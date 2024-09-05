<script setup lang="ts">
import TabItemTrigger from "@/components/item-trigger.vue";
import {ref, computed} from "vue";
import About from "@/components/about.vue";
const props = defineProps<{
  activeTab: string
}>()
const emits = defineEmits(['update:activeTab'])
const activeTabComp = computed({
  get: () => props.activeTab || 'modelValue',
  set: (m) => {
    emits('update:activeTab', m)
  }
});
const minimumOrder =ref()
const isOnAbout= ref(false)
</script>

<template>
  <div class="mt-8 flex justify-between border-b border-[#043D7F]  max-w-[1374px] w-full mx-auto px-8 border-box ">
    <div class="flex gap-8">
      <tab-item-trigger @click="isOnAbout=false" v-model="activeTabComp" value="modelValue">
        Cardápio
      </tab-item-trigger>
      <tab-item-trigger @click="isOnAbout=true" v-model="activeTabComp" value="finished">
        Sobre
      </tab-item-trigger>
    </div>
    <div class="flex justify-end">
      <button> Pedido mínimo: R$</button>
      <input class=" w-10" v-model="minimumOrder" :placeholder="minimumOrder">
    </div>
  </div>
  <div class="flex justify-end">
    <about v-if="isOnAbout" @close="isOnAbout=false" v-model="activeTabComp" value="modelValue"></about>
  </div>

</template>

<style scoped>

</style>