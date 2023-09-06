<script setup>
import { ref, watch, defineProps, defineEmits } from 'vue';
import TableRow from './TableRow.vue';

const props = defineProps({
    emailsObj: Object,
})

const emails = ref({});

watch(() => props.emailsObj, (emailsObj) => {
    emails.value = emailsObj;
})

defineEmits(['viewEmail'])
</script>
<template>
    <section class="mx-auto min-w-[30rem]">
        <div class="flex flex-col">
            <div class="overflow-x-auto">
                <div class="inline-block min-w-full align-middle ">
                    <div class="overflow-hidden border-2 border-zinc-400 rounded-lg">
                        <table class="divide-y divide-zinc-400">
                            <thead class="bg-green-500">
                                <tr>
                                    <th scope="col"
                                        class="px-4 py-3.5 text-sm font-bold text-left rtl:text-right text-slate-100">
                                        Subject
                                    </th>

                                    <th scope="col"
                                        class="px-4 py-3.5 text-sm font-bold text-left rtl:text-right text-slate-100">
                                        From
                                    </th>

                                    <th scope="col"
                                        class="px-4 py-3.5 text-sm font-bold text-left rtl:text-right text-slate-100">
                                        To
                                    </th>
                                </tr>
                            </thead>
                            <tbody class=" divide-y divide-zinc-400 bg-white">
                                <TableRow v-for="email in emails" :key="email._id"
                                    @view-email="(id) => $emit('viewEmail', emails.find((email) => email._id == id))"
                                    :subject="email._source.subject" :body="email._source.body"
                                    :from-name="email._source.xfrom" :from-email="email._source.from"
                                    :to-names="email._source.xto" :to-emails="email._source.to" />
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>
