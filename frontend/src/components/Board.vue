<script setup lang="ts">
// Script kısmında değişiklik yok
import { ref } from "vue";
import draggable from "vuedraggable";
import type { Board } from "../stores/boardStore";
import { useBoardStore } from "../stores/boardStore";
import TaskItem from "./TaskItem.vue";

const props = defineProps<{ board: Board }>();
const boardStore = useBoardStore();
const newTaskTitle = ref("");

const handleAddTask = () => {
  boardStore.addTask(props.board.id, newTaskTitle.value);
  newTaskTitle.value = "";
};

const onTaskDrop = (event: any) => {
  if (event.added) {
    const addedTask = event.added.element;
    addedTask.board_id = props.board.id;
    boardStore.updateTask(addedTask);
  }
};
</script>

<template>
  <div
    class="flex flex-col flex-shrink-0 w-80 bg-slate-200 rounded-xl shadow-md p-3"
  >
    <!-- Başlık rengini mora çevirelim -->
    <h2 class="font-bold text-lg mb-4 text-purple-700 px-1">
      {{ board.title }}
    </h2>

    <draggable
      v-model="board.tasks"
      group="tasks"
      item-key="id"
      class="space-y-3 min-h-[5rem] overflow-y-auto pr-1"
      @change="onTaskDrop"
    >
      <template #item="{ element }">
        <TaskItem :task="element" :boardId="board.id" />
      </template>
    </draggable>

    <form @submit.prevent="handleAddTask" class="mt-4">
      <textarea
        v-model="newTaskTitle"
        rows="2"
        placeholder="+ Yeni bir görev ekle"
        class="w-full p-2 bg-transparent text-slate-700 placeholder-slate-500 border-none rounded-lg focus:ring-2 focus:ring-purple-500 focus:bg-white hover:bg-slate-300 transition-colors duration-200"
        @keydown.enter.prevent="handleAddTask"
      ></textarea>
    </form>
  </div>
</template>
