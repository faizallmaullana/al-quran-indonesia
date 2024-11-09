<script setup lang="ts">
import Title from "@/components/wrappers/TitleCard.vue";
import { axios } from "@/axios/config";
import { onMounted, reactive } from "vue";
import { useRoute } from "vue-router";  // Import useRoute to get route parameters

interface Surah {
  aya_id: number;
  aya_number: number;
  aya_text: string;
  sura_id: number;
  juz_id: number;
  page_number: number;
  translation_aya_text: string;
}

interface SurahMeta {
  id: number;
  surat_name: string;
  surat_text: string;
  surat_terjemahan: string;
  count_ayat: number;
}


const state = reactive({
  surahs: [] as Surah[],
  surahMeta: {} as SurahMeta,
});

const route = useRoute();

onMounted(() => {
  getSurah();
});

function getSurah() {
  const surahId = route.params.surahid;
  axios.get(`/surah/${surahId}`)
    .then((res) => {
      state.surahs = res.data.data;
      state.surahMeta = res.data.meta
    })
    .catch((error) => {
      console.error("Error fetching surahs:", error);
    });
}

function cleanTranslationText(text: string): string {
  return text.replace(/<p>/g, '').replace(/<\/p>/g, '');
}
</script>


<template>
  <main class="bg-amber-50 flex flex-col gap-6">
    <Title :title="state.surahMeta.surat_text" :surah="state.surahMeta.surat_name" :terjemahan="state.surahMeta.surat_terjemahan" :ayat="state.surahMeta.count_ayat.toString()" />

    <h5 class="text-center my-5 text-emerald-700 bg-emerald-50 py-6">بِسْمِ اللّٰهِ الرَّحْمٰنِ الرَّحِيْمِ</h5>

    <div class="flex flex-col gap-16">
      <div class="flex flex-col gap-6" v-for="(surah, index) in state.surahs" :key="index">
        <div>
          <h5 class="text-right">{{ surah.aya_text }}</h5>
        </div>
        <p><strong class="text-emerald-700">({{ surah.aya_number }})</strong> {{ cleanTranslationText(surah.translation_aya_text) }}</p>
      </div>
    </div>
  </main>
</template>