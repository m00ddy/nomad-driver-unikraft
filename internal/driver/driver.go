package driver

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/drivers/shared/eventer"
	"github.com/hashicorp/nomad/plugins/base"
	"github.com/hashicorp/nomad/plugins/drivers"
	"github.com/hashicorp/nomad/plugins/shared/hclspec"
)

const (
	PluginName    = "unikraft"
	PluginVersion = "v0.1.0"
)

var _ drivers.DriverPlugin = (*Driver)(nil)

type Driver struct {
	eventer        *eventer.Eventer
	logger         hclog.Logger
	ctx            context.Context
	signalShutdown context.CancelFunc
}

var pluginInfo = &base.PluginInfoResponse{
	Type:              base.PluginTypeDriver,
	PluginApiVersions: []string{drivers.ApiVersion010},
	PluginVersion:     PluginVersion,
	Name:              PluginName,
}

func New(log hclog.Logger) drivers.DriverPlugin {
	// establish context
	ctx, cancel := context.WithCancel(context.Background())
	logger := log.Named(PluginName)

	return &Driver{
		eventer:        eventer.NewEventer(ctx, log.Named("Event")),
		logger:         logger,
		ctx:            ctx,
		signalShutdown: cancel,
	}

}

// we want to satisfy BasePlugin and DriverPlugin

// BasePlugin
func (d *Driver) PluginInfo() (*base.PluginInfoResponse, error) {
	return pluginInfo, nil
}

func (d *Driver) ConfigSchema() (*hclspec.Spec, error) {
	return nil, fmt.Errorf("not Implemented")
}

func (d *Driver) SetConfig(_ *base.Config) error {
	return fmt.Errorf("not implemented")
}

// DriverPlugin
func (d *Driver) TaskConfigSchema() (*hclspec.Spec, error) {
	return nil, fmt.Errorf("not implemented")
}

func (d *Driver) Capabilities() (*drivers.Capabilities, error) {
	return nil, fmt.Errorf("not implemented")
}

func (d *Driver) Fingerprint(_ context.Context) (<-chan *drivers.Fingerprint, error) {
	return nil, fmt.Errorf("not implemented")
}

func (d *Driver) RecoverTask(_ *drivers.TaskHandle) error {
	return fmt.Errorf("not implemented")
}

func (d *Driver) StartTask(_ *drivers.TaskConfig) (*drivers.TaskHandle, *drivers.DriverNetwork, error) {
	return nil, nil, fmt.Errorf("not implemented")
}

func (d *Driver) WaitTask(_ context.Context, _ string) (<-chan *drivers.ExitResult, error) {
	return nil, fmt.Errorf("not implemented")
}

func (d *Driver) StopTask(_ string, _ time.Duration, _ string) error {
	return fmt.Errorf("not implemented")
}

func (d *Driver) DestroyTask(_ string, _ bool) error {
	return fmt.Errorf("not implemented")
}

func (d *Driver) InspectTask(_ string) (*drivers.TaskStatus, error) {
	return nil, fmt.Errorf("not implemented")
}

func (d *Driver) TaskStats(_ context.Context, _ string, _ time.Duration) (<-chan *drivers.TaskResourceUsage, error) {
	return nil, fmt.Errorf("not implemented")
}

func (d *Driver) TaskEvents(_ context.Context) (<-chan *drivers.TaskEvent, error) {
	return nil, fmt.Errorf("not implemented")
}

func (d *Driver) SignalTask(_ string, _ string) error {
	return fmt.Errorf("not implemented")
}

func (d *Driver) ExecTask(_ string, _ []string, _ time.Duration) (*drivers.ExecTaskResult, error) {
	return nil, fmt.Errorf("not implemented")
}
