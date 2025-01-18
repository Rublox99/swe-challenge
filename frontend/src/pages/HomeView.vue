<script setup lang="ts">
import TopButtons from '../components/TopButtons.vue';
import Footer from '../components/Footer.vue';

import { userSession } from '@/services/sessionService';
import { ZincService } from '@/services/zincService';
import { onMounted, ref, toRaw } from 'vue';
import type { Hit, Hits } from '@/interfaces/Zinc';

var emails = ref<Hit[]>([])
var areEmailsLoading = ref(false)
var selectedEmail = ref<Hit>()
var isSelectedEmailLoading = ref(false)

var searchTerm = ref('')
var startDate = ref('')
var endDate = ref('')
var searchType = ref<'term' | 'id'>('term')

var page = ref(1)
var batchSize = ref(10)
var pageSizes = ref([5, 10, 15, 20])
var pagesAmount = ref(1)

const errorMessage = ref('')
const errorTitle = ref('')
const showAlert = ref(false)

onMounted(async () => {
    await fetchAllEmails()
})

const showErrorNotif = async () => {
    errorTitle.value = "Error: Emails"
    errorMessage.value = "An error occurred while fetching the filtered emails. Please, try again."

    showAlert.value = true
    await new Promise((resolve) => setTimeout(resolve, 2500))
    showAlert.value = false
}

const handleSearch = async (text: string, type: 'term' | 'id', start: string, end: string) => {
    searchTerm.value = text
    searchType.value = type

    startDate.value = start
    endDate.value = end

    await onParamChange()
}

const onParamChange = async () => {
    page.value = 1
    const isSearchEmpty = !searchTerm.value.trim()
    const isDateRangeEmpty = !startDate.value && !endDate.value

    switch (searchType.value) {
        case "term":
            isSearchEmpty && isDateRangeEmpty ? await fetchAllEmails() : await fetchFilteredEmails()
            break;

        case "id":
            if (searchTerm.value.length > 0)
                await fetchByID()
            else
                emails.value = []
            break;

        default:
            showErrorNotif()
            break;
    }
}

const onEmailClick = async (e: Hit) => {
    if (e._source['Message-Id'] === selectedEmail.value?._source['Message-Id']) {
        selectedEmail.value = undefined
        return
    }

    isSelectedEmailLoading.value = true

    await new Promise((resolve) => setTimeout(resolve, 1500))

    selectedEmail.value = e
    isSelectedEmailLoading.value = false
}

const fetchFilteredEmails = async () => {
    const pg = page.value
    const size = batchSize.value
    const from = (pg * size) - size

    areEmailsLoading.value = true

    try {
        await new Promise((resolve) => setTimeout(resolve, 1000))

        const { hits } = await ZincService.GetFilteredEmails(pg === 1 ? 0 : from, size, startDate.value, endDate.value, searchTerm.value)
        emails.value = hits.hits

        computeTotalValues(hits)
    } catch (error) {
        showErrorNotif()
    } finally {
        areEmailsLoading.value = false
    }
}

const fetchByID = async () => {
    areEmailsLoading.value = true

    page.value = 1
    pagesAmount.value = 1

    try {
        await new Promise((resolve) => setTimeout(resolve, 1000))

        const response = [await ZincService.GetEmailById(searchTerm.value)]

        if (response.length === 1 && response[0].error)
            emails.value = []
        else
            emails.value = response
    } catch (error) {
        showErrorNotif()
    } finally {
        areEmailsLoading.value = false
    }
}

const fetchAllEmails = async () => {
    const pg = page.value
    const size = batchSize.value
    const from = (pg * size) - size

    areEmailsLoading.value = true

    try {
        await new Promise((resolve) => setTimeout(resolve, 1000))

        const { hits } = await ZincService.GetAllEmails(pg === 1 ? 0 : from, size)
        emails.value = hits.hits

        computeTotalValues(hits)
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

const computeTotalValues = async (hits: Hits) => {
    const total = hits.total.value

    pagesAmount.value = Math.ceil(total / batchSize.value)
}
</script>

<template>
    <main class="relative flex flex-col gap-3 pb-2.5 md:pb-0 w-full h-full mx-auto rounded-md xl:max-w-[1650px]" :class="{
        'md:h-full': emails.length <= 11,
        'md:h-[92%]': emails.length > 11
    }">
        <TopButtons @onSearchEmit="handleSearch"></TopButtons>

        <section class="flex flex-col w-full sm:h-auto gap-2 md:pb-0.5 md:pt-1 md:flex-row" :class="{
            'md:h-full': emails.length <= 11,
            'md:h-[85.5%]': emails.length > 11
        }">
            <!-- Mails List -->
            <div class="w-full h-[75vh] rounded-md p-4 !md:max-w-[750px] md:h-full md:w-[60%] xs:border xs:border-gray-600 md:shadow-md"
                :class="{
                    'bg-gray-300': userSession.currentTheme.value === 'light',
                    'bg-gray-900': userSession.currentTheme.value === 'dark'
                }">

                <!-- Loader -->
                <div v-if="areEmailsLoading" class="flex items-center justify-center h-full">
                    <v-progress-circular color="blue-grey" model-value="100" indeterminate></v-progress-circular>
                </div>

                <!-- No emails -->
                <div v-else-if="emails.length === 0" class="flex items-center justify-center h-full">
                    <span class="animate-fade">{{ $t('home.placeholders.no_emails') }}</span>
                </div>

                <!-- Emails -->
                <div v-else-if="emails.length > 0"
                    class="w-full h-full flex flex-col justify-between gap-[10px] animate-fade">
                    <!-- Headers -->
                    <header class="flex justify-between gap-3 text-lg font-bold " :class="{
                        'text-gray-800': userSession.currentTheme.value === 'light',
                        'text-gray-400': userSession.currentTheme.value === 'dark'
                    }">
                        <span class="w-1/4">{{ $t('home.table.headers.subject') }}</span>
                        <span class="w-1/4 text-center">{{ $t('home.table.headers.date') }}</span>
                        <span class="w-1/4 text-center">{{ $t('home.table.headers.from') }}</span>
                        <span class="w-1/4 text-right">{{ $t('home.table.headers.to') }}</span>
                    </header>

                    <!-- Table -->
                    <section
                        class="flex-auto w-full h-full flex flex-col gap-[20px] animate-fade p-0 overflow-y-auto max-h-[calc(78.5%)]">
                        <div class="flex justify-between gap-5 py-1 transition-all rounded-md cursor-pointer hover:px-4"
                            v-on:click="onEmailClick(email)" :class="{
                                'hover:bg-primary-300 hover:text-white text-light-contrast': userSession.currentTheme.value === 'light',
                                'hover:bg-primary-600 hover:text-white text-dark-contrast': userSession.currentTheme.value === 'dark',
                                'bg-primary-600 px-2 py-1 text-white': email._source['Message-Id'] === selectedEmail?._source['Message-Id'] && userSession.currentTheme.value === 'light',
                                'bg-primary-300 px-2 py-1 text-white': email._source['Message-Id'] === selectedEmail?._source['Message-Id'] && userSession.currentTheme.value === 'dark'
                            }" v-for="email in emails" :key="email._source['Message-Id']">
                            <span class="w-1/4 truncate" :title="email._source.Subject">{{ email._source.Subject.length
                                === 0 ? '*' :
                                email._source.Subject }}</span>
                            <span class="w-1/4 text-center truncate" :title="new Date(email._source.Date).toISOString().replace('T', ':').slice(0, 19).replace(/-/g, '/').replace(/:/, '-')">
                                {{
                                    new Date(email._source.Date).toISOString()
                                        .replace('T', ':')
                                        .slice(0, 10)
                                }}
                            </span>
                            <span class="w-1/4 text-left truncate"
                                :title="email._source.From.length === 0 ? '*' : email._source.From">{{
                                    email._source.From.length === 0 ? '*' :
                                        email._source.From }}</span>
                            <span class="w-1/4 text-right truncate"
                                :title="email._source.To.length === 0 ? '*' : email._source.To.join(', ')">{{
                                    email._source.To.length === 0 ?
                                        '*'
                                        : email._source.To.join(', ') }}</span>
                        </div>
                    </section>

                    <!-- Pagination -->
                    <footer class="flex flex-col justify-between w-full md:flex-row">
                        <v-pagination class="xs:w-full md:w-[300px] lg:w-[500px]" :disabled="areEmailsLoading"
                            v-model="page" v-on:update:model-value="fetchFilteredEmails" size="small"
                            :length="pagesAmount" rounded="circle"></v-pagination>

                        <div class="w-[100px] xs:mx-auto xs:mt-2 md:mx-0 md:mt-0 md:w-[150px]">
                            <v-select v-model="batchSize" v-on:update:model-value="onParamChange"
                                :label="$t('home.page_size')" :hide-details="true" density="compact" :items="pageSizes"
                                variant="outlined">
                            </v-select>
                        </div>
                    </footer>
                </div>
            </div>

            <!-- On Email Pick -->
            <div class="relative w-full h-[75vh] p-4 rounded-md md:h-full md:w-[40%] overflow-y-auto xs:border xs:border-gray-600 md:shadow-md"
                :class="{
                    'bg-gray-300': userSession.currentTheme.value === 'light',
                    'bg-gray-900': userSession.currentTheme.value === 'dark'
                }">
                <!-- Linear Loader -->
                <v-progress-linear :active="isSelectedEmailLoading" :indeterminate="isSelectedEmailLoading" absolute
                    bottom></v-progress-linear>

                <div v-if="!selectedEmail" class="flex items-center justify-center w-full h-full">
                    <span :class="{
                        'text-gray-800': userSession.currentTheme.value === 'light',
                        'text-gray-600': userSession.currentTheme.value === 'dark'
                    }">
                        {{ $t('home.placeholders.pick_email') }}
                    </span>
                </div>

                <div v-else class="flex flex-col gap-2.5 animate-fade" :class="{
                    'text-light-contrast': userSession.currentTheme.value === 'light',
                    'text-dark-contrast': userSession.currentTheme.value === 'dark'
                }">
                    <span class="block w-full overflow-y-auto font-bold text-center max-h-20 no-scrollbar">
                        {{ selectedEmail._source.Subject.length === 0 ? '*' : selectedEmail._source.Subject }}
                    </span>

                    <span class="block w-full overflow-y-auto font-bold text-center max-h-20 no-scrollbar" :class="{
                        'text-gray-800': userSession.currentTheme.value === 'light',
                        'text-gray-600': userSession.currentTheme.value === 'dark'
                    }">
                        {{
                            new Date(selectedEmail._source.Date).toISOString()
                                .replace('T', ':')
                                .slice(0, 19)
                                .replace(/-/g, '/')
                                .replace(/:/, '-')
                        }}
                    </span>

                    <div class="w-full flex flex-wrap gap-2.5">
                        <span class="text-sm font-bold flex-0-0">{{ $t('home.email_content.from') }}:</span>
                        <span class="flex-auto text-sm font-normal" :class="{
                            'text-gray-800': userSession.currentTheme.value === 'light',
                            'text-gray-600': userSession.currentTheme.value === 'dark'
                        }">{{ selectedEmail._source.From }}</span>
                    </div>

                    <div class="w-full flex flex-wrap gap-2.5">
                        <span class="text-sm font-bold flex-0-0">{{ $t('home.email_content.to') }}:</span>
                        <span class="flex-auto text-sm font-normal" :class="{
                            'text-gray-800': userSession.currentTheme.value === 'light',
                            'text-gray-600': userSession.currentTheme.value === 'dark'
                        }">{{ selectedEmail._source.To.join(', ') }}</span>
                    </div>

                    <div class="w-full flex flex-wrap gap-2.5">
                        <span class="text-sm font-bold flex-0-0">CC:</span>
                        <span class="flex-auto text-sm font-normal" :class="{
                            'text-gray-800': userSession.currentTheme.value === 'light',
                            'text-gray-600': userSession.currentTheme.value === 'dark'
                        }">{{ selectedEmail._source.CC.join(', ') }}</span>
                    </div>

                    <div class="w-full flex flex-wrap gap-2.5">
                        <span class="text-sm font-bold flex-0-0">Content-Type:</span>
                        <span class="flex-auto text-sm font-normal" :class="{
                            'text-gray-800': userSession.currentTheme.value === 'light',
                            'text-gray-600': userSession.currentTheme.value === 'dark'
                        }">{{ selectedEmail._source['Content-Type'] }}</span>
                    </div>

                    <div class="w-full flex flex-wrap gap-2.5">
                        <span class="text-sm font-bold flex-0-0">ID:</span>
                        <span class="flex-auto text-sm font-normal" :class="{
                            'text-gray-800': userSession.currentTheme.value === 'light',
                            'text-gray-600': userSession.currentTheme.value === 'dark'
                        }">{{ selectedEmail._id }}</span>
                    </div>

                    <!-- Divisor -->
                    <div class="w-full h-[1.5px] bg-gray-800 my-3"></div>

                    <div class="w-full max-h-[350px] !md:max-h-200px">
                        <span class="text-left pb-3.5 font-light">{{ selectedEmail._source.Body }}</span>
                    </div>
                </div>

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
