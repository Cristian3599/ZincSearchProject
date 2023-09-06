<script setup>
import { onMounted, ref, watch, defineProps } from 'vue';
import MailTable from './MailTable.vue';
import MailDetail from './MailDetail.vue'

const props = defineProps({
  query: String,
});

const parentContainer = ref(null);
const scrollContainer = ref(null);
const emailDetail = ref({});
const emailsJson = ref({});

watch(() => props.query, async (query) => {
  emailsJson.value = await fetchEmails(query);
});

const scrollToEmailsList = () => {
  window.scrollTo({
    top: 0,
    behavior: 'smooth'
  })
  parentContainer.value.scrollTo({
    left: 0,
    behavior: 'smooth'
  });
};

const scrollToEmailDetail = () => {
  window.scrollTo({
    top: 0,
    behavior: 'smooth'
  })
  parentContainer.value.scrollTo({
    left: scrollContainer.value.offsetWidth,
    behavior: 'smooth'
  });
};

const viewEmail = (emailObj) => {
  scrollToEmailDetail();
  emailDetail.value = emailObj;
}

onMounted(async () => {
  emailsJson.value = await fetchEmails("");
})

const fetchEmails = async (query) => {
  const reqUrl = query == ""
    ? 'http://localhost:3000/emails'
    : `http://localhost:3000/emails/${query}`
  const emailsReq = await fetch(reqUrl);
  const reqResult = await emailsReq.json();
  return reqResult.hits.hits;
};
</script>
<style>
.parent-container::-webkit-scrollbar {
  display: none;
}
</style>
<template>
  <div class="parent-container w-screen overflow-x-scroll" ref="parentContainer">
    <div class="flex w-max" ref="scrollContainer">
      <div class="w-screen px-5 pb-5">
        <MailTable @view-email="viewEmail" :emails-obj="emailsJson"></MailTable>
      </div>
      <div class="w-screen px-5 pb-5">
        <MailDetail :email="emailDetail" @back-to-list="scrollToEmailsList"></MailDetail>
      </div>
    </div>
  </div>
</template>
