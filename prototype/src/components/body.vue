<script setup lang="ts">

const isEmpty = ref(true)
const activeTab = ref("")
import {onMounted, ref} from "vue"
import ItemsComponent from "@/components/items-component.vue";
const input = ref("")
const saida = ref([] as Array<{text:string, id:string}>)
onMounted(() => {
  get()
})
const get = async () => {


  // saida.value = response.data

}
</script>

<template>
  <div class="w-screen flex flex-row justify-center space-x-8  mt-4">

    <div class="flex flex-row">

      <div class=" flex content-between">
        <div class="box-border h-[390px] w-[375px] bg-[#FFFFFF] shadow-lg rounded-[40px] grid place-items-center" >
          <span class="justify-self-start ms-1 font-bold">Os maiores sucessos:</span>
          <img class="w-[80%] h-[80%]" src="@/assets/vecteezy_fast-food-meal-png-with-ai-generated_25064792.png">
          <div class="flex flex-col">
            <span class="uppercase">combo</span>
            <span> R$: 50,00</span>
            <button :hidden="false" @click="" class=" w-[105px] h-[35px] bg-amber-50 rounded-[8px] mt-4 "> edit </button>
          </div>

        </div>
      </div>

    </div>

    <div class="flex flex-row space-y-1">
      <items-component></items-component>
    </div>
  </div>
</template>

<style scoped>

</style>