package store

func (zk *ZkStore) UpdateCurrentTask(appId, slotId string, task *Task) error {
	appStore, found := zk.Apps[appId]
	if !found {
		return ErrAppNotFound
	}

	_, found = appStore.Slots[slotId]
	if !found {
		return ErrSlotNotFound
	}

	op := &StoreOp{
		Op:      OP_UPDATE,
		Entity:  ENTITY_CURRENT_TASK,
		Param1:  appId,
		Param2:  slotId,
		Payload: task,
	}

	return zk.Apply(op)
}

func (zk *ZkStore) ListTaskHistory(appId, slotId string) []*Task {
	appStore, found := zk.Apps[appId]
	if !found {
		return nil
	}

	slotStore, found := appStore.Slots[slotId]
	if !found {
		return nil
	}

	return slotStore.TaskHistory
}
