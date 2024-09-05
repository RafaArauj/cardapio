<script setup lang="ts">
import {useCounter} from '@/data/useCounter';
import TabItemTrigger from '@/components/item-trigger.vue';
import RegisterItems from '@/components/register-items.vue';
import { ref } from 'vue';

const {itemCount} = useCounter();
const isOnRegister = ref(false);
console.log(itemCount)
</script>

<template>
  <div class="mt-2 flex justify-between max-w-[1374px] w-full mx-auto px-8 border-box ">
    <div class="flex gap-8 m">
      <span class=" flex justify-center items-center font-medium ">
        {{ itemCount }} Cadastrados
      </span>

      <button @click="isOnRegister=true" class=" flex flex-row justify-center items-center bg-[#043D7F] rounded-[30px] text-white uppercase h-[40px] w-[180px]">
        <span class="material-symbols-outlined">
          add
        </span>
        <span>Cadastrar</span>
      </button>
    </div>
  </div>
  <div class=" flex items-center justify-center">
    <register-items v-if="isOnRegister" @close="isOnRegister=false" ></register-items>
  </div>
</template>
