<script setup lang="ts">
const emits = defineEmits(["close"])
const isClosed = ()=>{
  emits("close")
}

import toggle from "@/components/register-status.vue"
</script>

<template>
  <div class=" bg-gray-100 p-6 rounded-lg shadow-lg m-2 absolute ">
    <div class="flex justify-end items-center ">
      <button class="material-symbols-outlined" @click="isClosed">
        close
      </button>
    </div>
    <div class=" flex flex-col justify-center  items-center space-x-4 space-y-4 ">
    <button class="text-[100px] text-[#0078FC] material-symbols-outlined">
image
</button>
      <div class="flex justify-center flex-col items-center space-x-4 space-y-4">
        <input
            type="text"
            placeholder="Descrição"
            class="w-full text-center text-gray-800 font-semibold bg-transparent  focus:outline-none focus:ring-0"
        />


        <input
            type="text"
            placeholder="Preço"
            class="w-full text-center text-gray-800 font-semibold bg-transparent focus:outline-none focus:ring-0"
        />

        <button class="w-full bg-[#043D7F] text-white uppercase py-2  rounded-full hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          salvar
        </button>
      </div>
    </div>
  </div>

</template>

<style scoped>

</style>