<script setup lang="ts">
// Script kısmında değişiklik yok
import { onMounted, ref } from "vue";
import Board from "./components/Board.vue";
import TaskDetailModal from "./components/TaskDetailModal.vue";
import { useBoardStore } from "./stores/boardStore";

const boardStore = useBoardStore();
const newBoardTitle = ref("");

onMounted(() => {
  boardStore.fetchBoards();
});

const handleAddBoard = () => {
  boardStore.addBoard(newBoardTitle.value);
  newBoardTitle.value = "";
};
</script>

<template>
  <div class="min-h-screen bg-slate-100 font-sans p-4 sm:p-6">
    <TaskDetailModal v-if="boardStore.selectedTask" />
    <div>
      <header class="mb-6">
        <!-- Başlık rengini mora çevirelim -->
        <h1 class="text-4xl font-bold text-purple-800">Proje Panosu</h1>
      </header>

      <div class="flex overflow-x-auto space-x-4 pb-4">
        <Board
          v-for="board in boardStore.boards"
          :key="board.id"
          :board="board"
        />

        <div class="flex-shrink-0 w-80">
          <form @submit.prevent="handleAddBoard" class="p-2">
            <input
              v-model="newBoardTitle"
              type="text"
              placeholder="+ Yeni bir liste ekle"
              class="w-full bg-slate-200 hover:bg-slate-300 transition-colors duration-200 text-slate-700 placeholder-slate-500 p-3 rounded-lg border-none focus:ring-2 focus:ring-purple-500 focus:bg-white"
            />
          </form>
        </div>
      </div>
    </div>
  </div>
</template>
