<script setup lang="ts">
import type { Task } from "../stores/boardStore";
import { useBoardStore } from "../stores/boardStore";
// Heroicons'u import edelim
import { DocumentTextIcon, TrashIcon } from "@heroicons/vue/24/outline";

const props = defineProps<{ task: Task; boardId: string }>();
const boardStore = useBoardStore();

const openModal = () => {
  boardStore.selectTask(props.task);
};
</script>
<template>
  <div
    @click="openModal"
    class="group relative flex flex-col p-3 bg-white rounded-lg shadow-sm hover:shadow-lg transition-shadow duration-200 cursor-pointer"
  >
    <div class="flex items-start justify-between">
      <p class="text-base text-slate-800 break-words">{{ task.title }}</p>
      <button
        @click.stop="boardStore.deleteTask(boardId, task.id)"
        class="absolute top-2 right-2 p-1 rounded-full text-gray-400 hover:text-red-500 opacity-0 group-hover:opacity-100 transition-opacity"
      >
        <!-- Eski SVG yerine Heroicon component'i -->
        <TrashIcon class="h-5 w-5" />
      </button>
    </div>

    <div v-if="task.content" class="mt-2">
      <!-- Eski SVG yerine Heroicon component'i ve daha küçük boyut -->
      <DocumentTextIcon class="h-4 w-4 text-gray-400" />
    </div>
  </div>
</template>
