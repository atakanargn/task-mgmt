<script setup lang="ts">
import { ref, watch } from "vue";
import type { Task } from "../stores/boardStore";
import { useBoardStore } from "../stores/boardStore";
// Heroicons'u import edelim
import { CheckBadgeIcon, XMarkIcon } from "@heroicons/vue/24/solid";

const boardStore = useBoardStore();
const taskData = ref<Task | null>(null);

watch(
  () => boardStore.selectedTask,
  (newTask) => {
    if (newTask) {
      taskData.value = { ...newTask };
    } else {
      taskData.value = null;
    }
  },
  { immediate: true }
);

const saveAndClose = () => {
  if (taskData.value && taskData.value.id) {
    boardStore.updateTask(taskData.value);
  }
  boardStore.clearSelectedTask();
};

const closeModal = () => {
  boardStore.clearSelectedTask();
};
</script>

<template>
  <div
    @click.self="closeModal"
    class="fixed inset-0 bg-black bg-opacity-60 z-40 flex items-center justify-center p-4"
  >
    <div
      v-if="taskData"
      class="bg-slate-50 rounded-lg shadow-xl w-full max-w-2xl p-6 z-50"
    >
      <div class="flex items-start justify-between mb-4">
        <div class="flex items-center gap-2">
          <!-- Yeşil renkli bir ikon ekleyelim -->
          <CheckBadgeIcon class="h-8 w-8 text-green-500" />
          <input
            v-model="taskData.title"
            type="text"
            class="text-2xl font-bold w-full border-none bg-transparent focus:ring-0 focus:outline-none text-slate-800 p-0"
          />
        </div>
        <button @click="closeModal" class="p-1 rounded-full hover:bg-slate-200">
          <!-- Eski SVG yerine Heroicon component'i -->
          <XMarkIcon class="h-6 w-6 text-slate-500" />
        </button>
      </div>

      <h3 class="font-semibold text-slate-600 mb-2 mt-6">Açıklama</h3>
      <textarea
        v-model="taskData.content"
        rows="8"
        placeholder="Görev için bir açıklama ekleyin..."
        class="w-full bg-slate-100 border-slate-200 rounded-md focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
      ></textarea>

      <div class="flex justify-end mt-6">
        <!-- Kaydet butonunu mora çevirelim -->
        <button
          @click="saveAndClose"
          class="px-6 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700 transition-colors"
        >
          Kaydet ve Kapat
        </button>
      </div>
    </div>
  </div>
</template>
