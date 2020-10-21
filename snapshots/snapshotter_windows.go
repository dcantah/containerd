package snapshots

// WithScratchSpace sets a label to be utilized by a windows snapshotter. The label
// signifies that this snapshot should create a scratch space vhdx to be used as the
// r/w layer for a container.
func WithScratchSpace() Opt {
	return func(info *Info) error {
		if info.Labels == nil {
			info.Labels = make(map[string]string)
		}
		info.Labels["containerd.io/snapshot/io.microsoft.container.scratch-space"] = ""
		return nil
	}
}
