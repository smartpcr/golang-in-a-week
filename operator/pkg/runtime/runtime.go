package runtime

import (
	"operator/pkg/subscription"
	"sync"
)

func RunLoop(subscriptions []subscription.ISubscription) error {
	var wg sync.WaitGroup
	for _, sub := range subscriptions {
		wg.Add(1)
		go func(subscription subscription.ISubscription) {
			defer wg.Done()
			watchInterface, err := subscription.Subscribe()
			if err != nil {
				return
			}
			for {
				select {
				case event, ok := <-watchInterface.ResultChan():
					if !ok {
						return
					}
					err := subscription.Reconcile(event.Object, event.Type)
					if err != nil {
						return
					}
				case <-subscription.IsComplete():
					return
				}
			}
		}(sub)
	}
	wg.Wait()
	return nil
}
