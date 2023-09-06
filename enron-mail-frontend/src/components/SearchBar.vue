<script setup>
import { reactive, ref, defineEmits } from 'vue';

const query = reactive({ query: "" });
const lastQuery = ref("");

const emit = defineEmits(['changeQuery'])

const handleSubmit = () => {
  if (lastQuery.value != query.query) {
    lastQuery.value = query.query;
    emit('changeQuery', query.query);
  }
}
</script>

<template>
  <form @submit.prevent="handleSubmit">
    <div class="flex flex-row items-center w-full p-5">
      <div class="border-solid border-green-400 border-2 box-border h-10 rounded-2xl py-1 px-3 w-full mr-2 bg-green-200">
        <input type="text" class="focus:outline-none w-full bg-green-200 text-black" v-model="query.query" />
      </div>
      <button type="submit" class="bg-slate-800 hover:border-zinc-400 border-slate-800 border-2 box-border h-10 text-slate-100 flex flex-row items-center px-3 py-1.5 rounded-3xl ml-2">
        <font-awesome-icon icon="fa-solid fa-magnifying-glass" style="color: #f6f4f4;" />
      </button>

    </div>
  </form>
</template>
