<script setup lang="ts">
import { ref } from "vue";
import { useTaskStore } from "../stores/taskStore";

const taskStore = useTaskStore();
const newTaskTitle = ref("");

const handleSubmit = () => {
  taskStore.addTask(newTaskTitle.value);
  newTaskTitle.value = "";
};
</script>

<template>
  <form @submit.prevent="handleSubmit" class="flex gap-2 mb-4">
    <input
      type="text"
      v-model="newTaskTitle"
      placeholder="Task title"
      class="flex-grow rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
    />
    <button
      type="submit"
      class="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
    >
      Add
    </button>
  </form>
</template>
