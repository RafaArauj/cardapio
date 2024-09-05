<script setup lang="ts">
import ClubVoaBody from './components/body.vue';
import ClubVoaBodyListingHeader from "@/components/body-listing-header.vue";
import {ref} from "vue";
import RegisterStatus from "@/components/register-status.vue";
const activeTab = ref("")

</script>

<template>
  <main class=" h-screen w-screen flex flex-col">
    <div class="flex justify-center bg-[#0078FC] h-[150px] w-full ">
      <span class="mt-8" >
        <img src="@/assets/restaurante.svg">
      </span>
    </div>
    <club-voa-body-listing-header v-model:active-tab="activeTab"></club-voa-body-listing-header>
    <div class="flex justify-center items-center">

    </div>
    <register-status></register-status>
    <club-voa-body></club-voa-body>

  </main>
</template>

<style scoped>

</style>
