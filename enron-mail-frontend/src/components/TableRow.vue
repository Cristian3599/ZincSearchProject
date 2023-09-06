<script setup>
import { computed, ref, getCurrentInstance, defineProps, defineEmits } from 'vue';
import ProfileRow from './ProfileRow.vue';

const component = getCurrentInstance();

const props = defineProps({
  subject: String,
  body: String,
  fromName: String,
  fromEmail: String,
  toNames: Array,
  toEmails: Array,
})

const subjectRef = ref(props.subject)
const bodyRef = ref(props.body)

const getSubject = computed(() =>
  subjectRef.value.length != 0
    ? subjectRef.value
    : `Body: ${bodyRef.value.slice(0, 20)}`)

defineEmits(['viewEmail'])

</script>

<template>
  <tr class="cursor-pointer w-full" @click="$emit('viewEmail', component.vnode.key)">
    <td class="p-4 text-md text-black">
      <h1>
        {{ getSubject }}
      </h1>
    </td>
    <td class="p-2 lg:p-4 text-sm max-w-[14rem] lg:max-w-md w-2/12 truncate">
      <ProfileRow :names="[fromName]" :emails="[fromEmail]"></ProfileRow>
    </td>
    <td class="p-2 lg:p-4 text-sm max-w-[14rem] lg:max-w-md w-2/12">
      <ProfileRow :names="toNames" :emails="toEmails"></ProfileRow>
    </td>
  </tr>
</template>
