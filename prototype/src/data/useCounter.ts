// src/composables/useCounter.ts
import { computed } from 'vue';
import { listItems } from '@/data/items-list';

export function useCounter() {
    const itemCount = computed(() => listItems.length);

    return {
        itemCount,
    };
}
