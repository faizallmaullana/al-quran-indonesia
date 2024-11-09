<script setup lang="ts">
import Title from "@/components/wrappers/TitleCard.vue";
import { axios } from "@/axios/config";
import { onMounted, reactive } from "vue";
import router from '@/router/index'

interface Surah {
  id: number;
  surat_name: string;
  surat_text: string;
  surat_terjemahan: string;
  count_ayat: number;
}

const state = reactive({
  surahs: [] as Surah[],
});

onMounted(() => {
  getSurah();
});

function getSurah() {
  axios.get("/surah")
    .then((res) => {

      state.surahs = res.data.data;
    })
    .catch((error) => {
      console.error("Error fetching surahs:", error);
    });
}

function pushPath(path: string) {
  router.push(path)
}
</script>

<template>
  <main class="bg-amber-50 flex flex-col gap-6">
    <Title surah="Faizal Maulana" terjemahan="3IA04 | 50422497" title="" ayat="" />

    <div>
      <p class="font-semibold text-xl text-emerald-800">Surah</p>
    </div>

    <div class="flex flex-col gap-6">
      <div class="flex flex-col gap-6" v-for="(surah, index) in state.surahs" :key="index">
        <div class="flex justify-between items-end" @click.prevent="pushPath(`/surah/${surah.id}`)">
          <div class="flex gap-6 items-end">
            <p class="w-12 h-12 text-white bg-emerald-600 rounded-lg flex items-center justify-center">
              {{ surah.id }}</p>
            <div>
              <p class="text-gray-800 font-semibold text-lg">{{ surah.surat_name }}</p>
              <p class="text-gray-500 text-sm">{{ surah.surat_terjemahan }}</p>
            </div>
          </div>
          <div>
            <h5 class="text-emerald-800">{{ surah.surat_text }}</h5>
          </div>
        </div>
        <hr>
      </div>
    </div>
  </main>
</template>
