<script lang="ts">
import { defineComponent } from 'vue';
import { listItems } from '@/data/items-list';

export default defineComponent({
  methods: {
    listItems() {
      return listItems;
    }
  }
});
</script>

<template>
  <div class="p-6">
    <h2 class="text-xl font-bold mb-4">Hamburgers:</h2>
    <div class="grid grid-cols-1 sm:grid-cols-2  lg:grid-cols-3 gap-6">
      <div v-for="item in listItems()" class="px-1">
        <div class="box-border  bg-[#FFFFFF] shadow-lg rounded-[20px]">
          <div class="flex flex-col items-center justify-center m-8">
            <span>
              <img :src="item.offer[0].image" :alt="item.offer[0].name" class="w-32 h-32" />
            </span>
            <span class="mt-2 text-lg font-semibold">{{ item.offer[0].name }}</span>
            <span class="text-blue-600 font-bold">R$: {{ item.offer[0].price }}</span>
            <div class="flex flex-row items-center justify-center mt-2 space-x-4">
              <button @click="" class="bg-none rounded-[8px]">edit</button>
              <span @click="" class="material-symbols-outlined text-black cursor-pointer">
                delete_forever
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Você pode adicionar estilos personalizados aqui */
</style>
