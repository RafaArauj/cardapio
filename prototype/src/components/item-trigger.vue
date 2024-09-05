<script setup lang="ts">
import {computed} from "vue";

const props = defineProps<{
  value: string,
  modelValue: string
}>()
const emits = defineEmits(["update:modelValue"])
const classes = computed( ()=>{
  if(props.value == props.modelValue){
    return "border-[#043D7F]"
  }
  return "border-[#00000000]"
})
const onClick = () =>{
  emits("update:modelValue", props.value)
}
</script>

<template>
  <span @click="onClick" :class="['font-light text-[20px] pb-1 border-box border-b-[6px]', classes ]">
    <slot>

    </slot>
  </span>
</template>

<style scoped>

</style>