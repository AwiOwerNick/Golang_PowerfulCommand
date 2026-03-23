package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	ext     []string
	size    int64
	list    bool
	del     bool
	wLog    io.Writer
	archive string
	len     int
}

func main() {
	root := flag.String("root", ".", "Root directory to start")
	list := flag.Bool("list", false, "List files only")
	size := flag.Int64("size", 0, "Minimum file size")
	del := flag.Bool("del", false, "Delete files")
	log := flag.String("log", "", "Log deletes to this file")
	archive := flag.String("archive", "", "Archive directory")
	len := flag.Int("len", 0, "Minimum file name length")
	flag.Parse()
	var (
		f   = os.Stdout
		err error
	)

	if *log != "" {
		f, err = os.OpenFile(*log, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	c := config{
		ext:     flag.Args(),
		size:    *size,
		list:    *list,
		del:     *del,
		wLog:    f,
		archive: *archive,
		len:     *len,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	delLogger := log.New(cfg.wLog, "DELETED FILE: ", log.LstdFlags)

	return filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if len(cfg.ext) != 0 {
				ok := true
				for _, v := range cfg.ext {
					if !filterOut(path, v, cfg.size, info, cfg.len) {
						ok = false
					}
				}
				if ok {
					return nil
				}
			} else {
				if filterOut(path, "", cfg.size, info, cfg.len) {
					return nil
				}
			}

			if cfg.list {
				return listFile(path, out)
			}

			if cfg.archive != "" {
				if err := archiveFile(cfg.archive, root, path); err != nil {
					return err
				}
			}

			if cfg.del {
				return delFile(path, delLogger)
			}

			return listFile(path, out)

		})
}
