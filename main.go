package main

import (
	"context"
	_ "embed"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed main.wasm
var binary []byte

func main() {
	cfg := wazero.
		NewRuntimeConfig().
		WithDebugInfoEnabled(true).
		WithMemoryCapacityFromMax(true)

	ctx := context.Background()
	runtime := wazero.NewRuntimeWithConfig(ctx, cfg)

	wasi_snapshot_preview1.MustInstantiate(ctx, runtime)

	env := runtime.NewHostModuleBuilder("env")
	env = env.NewFunctionBuilder().WithFunc(func(int32) int32 { panic("__cxa_allocate_exception") }).Export("__cxa_allocate_exception")
	env = env.NewFunctionBuilder().WithFunc(func(int32, int32, int32) { panic("__cxa_throw") }).Export("__cxa_throw")
	if _, err := env.Instantiate(ctx); err != nil {
		return
	}

	compiled, err := runtime.CompileModule(ctx, binary)
	if err != nil {
		panic(err)
	}

	moduleConfig := wazero.
		NewModuleConfig().
		WithStdout(os.Stdout).
		WithStderr(os.Stderr).
		WithFSConfig(wazero.NewFSConfig().WithDirMount("/", "/"))

	m, err := runtime.InstantiateModule(ctx, compiled, moduleConfig)
	if err != nil {
		panic(err)
	}

	if _, err := m.ExportedFunction("do_debug").Call(ctx); err != nil {
		panic(err)
	}
}
