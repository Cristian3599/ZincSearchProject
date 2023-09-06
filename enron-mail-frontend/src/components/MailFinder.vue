<script setup>
  import { ref } from 'vue';

  import SearchBar from './SearchBar.vue';
  import MailViewer from './MailViewer.vue';

  const queryRef = ref("");

  const handleQuery = (query) => {
    queryRef.value = query;
  }
</script>
<template>
  <div class="bg-white">
    <SearchBar @change-query="handleQuery"></SearchBar>
    <MailViewer :query="queryRef"></MailViewer>
  </div>
</template>
