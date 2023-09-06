<script setup>
import { computed, ref, defineProps } from 'vue';

const props = defineProps({
  names: Array,
  emails: Array,
})

const namesRef = ref(props.names);
const emailsRef = ref(props.emails);

const getImgs = computed(() => {
  return `https://i.pravatar.cc/150?u=${emailsRef.value?.[0]}`
})

const removeTags = (stringWithTags) => stringWithTags?.replace(/<[^>]+>/g, "")

const getNames = computed(() => {
  return namesRef.value?.map((name) => removeTags(name));
})

const getEmails = computed(() => {
  return emailsRef.value?.map((email) => removeTags(email));
})

const getNamesList = computed(() => {
  return getNames.value?.slice(1,).join(" - ");
})

const getEmailsList = computed(() => {
  return getEmails.value?.slice(1,).join(" - ");
})
</script>

<template>
  <div class="flex items-center gap-x-3 w-full">
    <img class="hidden lg:block object-cover w-8 h-8 rounded-full" :src="getImgs" alt="">
    <div class="w-full">
      <div class="flex items-center w-10/12">
        <h2 class="text-sm font-medium text-black mr-2 whitespace-normal text-ellipsis">
          {{ getNames?.[0] }}
        </h2>
        <div class="popover popover-hover">
          <span class="badge popover-trigger" v-if="getNames?.length > 1">+{{ getNames?.length - 1 }}</span>
          <div class="popover-content popover-left-bottom whitespace-normal">
            <div class="p-4 text-sm">{{ getNamesList }}</div>
          </div>
        </div>
      </div>

      <div class="flex items-center">
        <p class="text-xs font-normal text-black mr-2 whitespace-normal text-ellipsis">
          {{ getEmails?.[0] }}
        </p>

        <div class="popover popover-hover">
          <span class="badge popover-trigger" v-if="getEmails?.length > 1">+{{ getEmails?.length - 1 }}</span>
          <div class="popover-content popover-left-bottom whitespace-normal">
            <div class="p-4 text-sm">{{ getEmailsList }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
