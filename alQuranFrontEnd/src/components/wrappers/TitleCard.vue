<template>
  <div>
    <div
      class="flex justify-between items-center relative p-6 rounded-xl text-white bg-gradient-to-r from-emerald-900 to-emerald-500 truncate">
      <div class="flex flex-col justify-between gap-6 z-10">
        <h2 class="text-2xl font-semibold">
          <span v-if="!title">Al Quran</span>
          <span v-else>{{ title }}</span>
        </h2>
        <div>
          <p class="font-semibold text-lg">{{ surah }}</p>
          <p class="text-sm">{{ terjemahan }}</p>
          <p class="text-sm" v-if="ayat">{{ ayat }} ayat</p>
        </div>
      </div>
      <div class="z-10">
        <img src="../imgs/quran.svg" alt="quran" class="w-[140px]">
      </div>
      <img src="../imgs/vector.svg" class="absolute bottom-[-20px] left-0 z-0 w-full">
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{ surah: string; terjemahan: string; title: string; ayat: string }>()
</script>
