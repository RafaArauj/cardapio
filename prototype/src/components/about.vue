<template>
  <div class="bg-white p-6 rounded-lg shadow-lg absolute">
    <div class="mb-4">
      <label for="description" class="block text-lg font-bold">Descrição:</label>
      <textarea
          id="description"
          rows="4"
          class="w-full mt-1 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Digite a descrição aqui..."
          v-model="description"
      ></textarea>
    </div>

    <div class="mb-4">
      <label for="openinghours" class="block text-lg font-bold">Horário de funcionamento:</label>
      <textarea
          id="openinghours"
          rows="8"
          class="w-full mt-1 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Digite os horários aqui..."
          v-model="openingHours"
      ></textarea>
    </div>

    <div class="mb-4">
      <label for="address" class="block text-lg font-bold">Endereço:</label>
      <textarea
          id="address"
          rows="3"
          class="w-full mt-1 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Digite o endereço aqui..."
      ></textarea>
    </div>

    <div class="mb-4">
      <label for="payment" class="block text-lg font-bold">Métodos de pagamento:</label>
      <input
          id="payment"
          type="text"
          class="w-full mt-1 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Digite os métodos de pagamento aqui..."
          v-model="payment"
      />
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';

export default {
  setup() {
    const description = ref('');
    const openingHours = ref('');
    const payment = ref('');

    return {
      description,
      openingHours,
      payment,
    };
  },
};
</script>
