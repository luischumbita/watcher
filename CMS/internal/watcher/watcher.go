// watcher.go
package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func setupWatcher(filePath string, callback func()) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Error al crear watcher:", err)
	}
	defer watcher.Close()

	if err := watcher.Add(filePath); err != nil {
		log.Fatal("Error al agregar archivo al watcher:", err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("Archivo modificado:", event.Name)
				callback()
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error del watcher:", err)
		}
	}
}
