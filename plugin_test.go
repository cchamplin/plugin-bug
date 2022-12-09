package plugin_bug

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/cchamplin/plugin-bug/plugin"
)

func TestPlugin_MemorySmall(t *testing.T) {
	ctx := context.Background()
	p, err := plugin.NewDataPasserPlugin(ctx, plugin.DataPasserPluginOption{})
	if err != nil {
		t.Error(err)
		return
	}
	defer p.Close(ctx)

	myPlugin, err := p.Load(ctx, "plugin.wasm")
	if err != nil {
		t.Error(err)
		return
	}

	entrySize := uint64(1000)
	response, err := myPlugin.PassData(context.Background(), plugin.PassDataRequest{
		Entries: entrySize,
	})
	if err != nil {
		t.Error(err)
		return
	}

	if len(response.Data) != int(entrySize) {
		t.Errorf("expected %d entries, got %d", entrySize, len(response.Data))
		return
	}

	log.Printf("Size of data being passed %d (%d bytes)", len(response.Data), response.SizeVT())

	for i := uint64(0); i < 20; i++ {
		log.Printf("calling plugin %d", i)
		start := time.Now()
		response, err = myPlugin.PassData(context.Background(), plugin.PassDataRequest{
			Sequence: i,
			Entries:  entrySize,
			Data:     response.Data,
		})
		log.Printf("call execution time: %s", time.Since(start))
		if err != nil {
			t.Error(err)
			return
		}

		if len(response.Data) != int(entrySize) {
			t.Errorf("expected %d entries, got %d", entrySize, len(response.Data))
			return
		}
	}
}

func TestPlugin_MemoryMedium(t *testing.T) {
	ctx := context.Background()
	p, err := plugin.NewDataPasserPlugin(ctx, plugin.DataPasserPluginOption{})
	if err != nil {
		t.Error(err)
		return
	}
	defer p.Close(ctx)

	myPlugin, err := p.Load(ctx, "plugin.wasm")
	if err != nil {
		t.Error(err)
		return
	}

	entrySize := uint64(5000)
	response, err := myPlugin.PassData(context.Background(), plugin.PassDataRequest{
		Entries: entrySize,
	})
	if err != nil {
		t.Error(err)
		return
	}

	if len(response.Data) != int(entrySize) {
		t.Errorf("expected %d entries, got %d", entrySize, len(response.Data))
	}

	log.Printf("Size of data being passed %d (%d bytes)", len(response.Data), response.SizeVT())

	for i := uint64(0); i < 20; i++ {
		log.Printf("calling plugin %d", i)
		start := time.Now()
		response, err = myPlugin.PassData(context.Background(), plugin.PassDataRequest{
			Sequence: i,
			Entries:  entrySize,
			Data:     response.Data,
		})
		log.Printf("call execution time: %s", time.Since(start))
		if err != nil {
			t.Error(err)
			return
		}

		if len(response.Data) != int(entrySize) {
			t.Errorf("expected %d entries, got %d", entrySize, len(response.Data))
			return
		}
	}
}

func TestPlugin_MemoryBig(t *testing.T) {
	ctx := context.Background()
	p, err := plugin.NewDataPasserPlugin(ctx, plugin.DataPasserPluginOption{})
	if err != nil {
		t.Error(err)
		return
	}
	defer p.Close(ctx)

	myPlugin, err := p.Load(ctx, "plugin.wasm")
	if err != nil {
		t.Error(err)
		return
	}

	entrySize := uint64(10000)
	response, err := myPlugin.PassData(context.Background(), plugin.PassDataRequest{
		Entries: entrySize,
	})
	if err != nil {
		t.Error(err)
		return
	}

	if len(response.Data) != int(entrySize) {
		t.Errorf("expected %d entries, got %d", entrySize, len(response.Data))
		return
	}

	log.Printf("Size of data being passed %d (%d bytes)", len(response.Data), response.SizeVT())

	for i := uint64(0); i < 20; i++ {
		log.Printf("calling plugin %d", i)
		start := time.Now()
		response, err = myPlugin.PassData(context.Background(), plugin.PassDataRequest{
			Sequence: i,
			Entries:  entrySize,
			Data:     response.Data,
		})
		log.Printf("call execution time: %s", time.Since(start))
		if err != nil {
			t.Error(err)
			return
		}

		if len(response.Data) != int(entrySize) {
			t.Errorf("expected %d entries, got %d", entrySize, len(response.Data))
			return
		}
	}
}

func TestPlugin_Error(t *testing.T) {
	ctx := context.Background()
	p, err := plugin.NewDataPasserPlugin(ctx, plugin.DataPasserPluginOption{})
	if err != nil {
		t.Error(err)
		return
	}
	defer p.Close(ctx)

	myPlugin, err := p.Load(ctx, "plugin.wasm")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = myPlugin.PassDataError(context.Background(), plugin.PassDataRequest{})
	if err == nil {
		t.Error("expected error")
		return
	}
}
