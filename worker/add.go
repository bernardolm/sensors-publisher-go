package worker

func (w *worker) Add(t ...taskFunc) {
	w.flows = append(w.flows, t)
}
