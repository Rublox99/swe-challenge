<script setup lang="ts">
import TopButtons from '../components/TopButtons.vue';
import Footer from '../components/Footer.vue';

import { userSession } from '@/services/sessionService';
import { ZincService } from '@/services/zincService';
import { onMounted, ref, toRaw } from 'vue';
import type { EmailSource } from '@/interfaces/Zinc';

var emails = ref<EmailSource[]>([])
var areEmailsLoading = ref(false)

var page = ref(1)
var batchSize = ref(10)
var pageSizes = ref([5, 10, 15, 20])

const errorMessage = ref('')
const errorTitle = ref('')
const showAlert = ref(false)

onMounted(() => {
    fetchAllEmails()
})

const onParamChange = async () => {
    page.value = 1
    await fetchAllEmails()
}

const fetchAllEmails = async () => {
    const pg = page.value
    const size = batchSize.value

    areEmailsLoading.value = true

    try {
        await new Promise((resolve) => setTimeout(resolve, 1000))

        const { hits } = await ZincService.GetAllEmails(pg === 1 ? 0 : (pg * size) - size, size)
        emails.value = hits.hits.map((hit) => hit._source)

        console.log(toRaw(emails.value))
    } catch (error) {
        errorTitle.value = "Error: Emails"
        errorMessage.value = "An error occurred while fetching the emails. Please, try again."

        showAlert.value = true
        await new Promise((resolve) => setTimeout(resolve, 2500))
        showAlert.value = false
    } finally {
        areEmailsLoading.value = false
    }
}
</script>

<template>
    <main class="relative flex flex-col gap-3 pb-2.5 md:pb-0 w-full h-full mx-auto rounded-md xl:max-w-[1500px]" :class="{
        'md:h-full': emails.length <= 11,
        'md:h-[92%]': emails.length > 11
    }">
        <TopButtons></TopButtons>

        <section class="flex flex-col w-full sm:h-auto gap-2 md:pb-0.5 md:pt-1 md:flex-row" :class="{
            'md:h-full': emails.length <= 11,
            'md:h-[85.5%]': emails.length > 11
        }">
            <div class="w-full h-[75vh] rounded-md p-4 !md:max-w-[750px] md:h-full md:w-[60%]" :class="{
                'bg-gray-300': userSession.currentTheme.value === 'light',
                'bg-gray-900': userSession.currentTheme.value === 'dark'
            }">

                <!-- Loader -->
                <div v-if="areEmailsLoading" class="flex items-center justify-center h-full">
                    <v-progress-circular color="blue-grey" model-value="100" indeterminate></v-progress-circular>
                </div>

                <!-- No emails -->
                <div v-else-if="emails.length === 0" class="flex items-center justify-center h-full">
                    <span>No emails were found</span>
                </div>

                <!-- Emails -->
                <div v-else-if="emails.length > 0"
                    class="w-full h-full flex flex-col justify-between gap-[10px] animate-fade">
                    <!-- Headers -->
                    <header class="flex justify-between gap-3 text-lg font-bold " :class="{
                        'text-gray-800': userSession.currentTheme.value === 'light',
                        'text-gray-400': userSession.currentTheme.value === 'dark'
                    }">
                        <span class="w-1/3">Subject</span>
                        <span class="w-1/3 text-center">From</span>
                        <span class="w-1/3 text-right">To</span>
                    </header>

                    <!-- Table -->
                    <section
                        class="flex-auto w-full h-full flex flex-col gap-[20px] animate-fade p-0 overflow-y-auto max-h-[calc(78.5%)]">
                        <div class="flex justify-between gap-5 py-1 transition-all rounded-md cursor-pointer hover:px-4"
                            :class="{
                                'hover:bg-primary-300 hover:text-white text-light-contrast': userSession.currentTheme.value === 'light',
                                'hover:bg-primary-500 hover:text-white text-dark-contrast': userSession.currentTheme.value === 'dark'
                            }" v-for="email in emails" :key="email['Message-Id']">
                            <span class="w-1/3 truncate" :title="email.Subject">{{ email.Subject }}</span>
                            <span class="w-1/3 text-center truncate" :title="email.From">{{ email.From }}</span>
                            <span class="w-1/3 text-right truncate" :title="email.To.join(', ')">{{ email.To.join(', ')
                                }}</span>
                        </div>
                    </section>

                    <!-- Pagination -->
                    <footer class="flex flex-col justify-between w-full md:flex-row">
                        <v-pagination :disabled="areEmailsLoading" v-model="page"
                            v-on:update:model-value="fetchAllEmails" size="small" length="4"
                            rounded="circle"></v-pagination>

                        <div class="w-[100px] mx-auto mt-3 md:mt-0 md:w-[150px] md:mx-0 md:ml-auto">
                            <v-select v-model="batchSize" v-on:update:model-value="onParamChange" label="Page Size"
                                hide-details :items="pageSizes" variant="outlined">
                            </v-select>
                        </div>
                    </footer>
                </div>
            </div>

            <div class="w-full h-[75vh] rounded-md md:h-full md:w-[40%]" :class="{
                'bg-gray-300': userSession.currentTheme.value === 'light',
                'bg-gray-900': userSession.currentTheme.value === 'dark'
            }">

            </div>
        </section>

        <!-- Alert Section -->
        <section v-if="showAlert"
            class="absolute w-full bottom-[10%] md:bottom-[5%] left-[50%] translate-x-[-50%] md:w-[400px] animate-fade">
            <v-alert :title="errorTitle" :text="errorMessage" type="error" variant="tonal">
            </v-alert>
        </section>

        <Footer></Footer>
    </main>
</template>
