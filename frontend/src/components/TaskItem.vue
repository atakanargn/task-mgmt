<script setup lang="ts">
import { useTaskStore } from "../stores/taskStore";

import type { Task } from "../stores/taskStore";

const props = defineProps<{
  task: Task;
}>();

const taskStore = useTaskStore();

const toggleCompleted = () => {
  const updatedTask = { ...props.task, completed: !props.task.completed };
  taskStore.updateTask(updatedTask);
};
</script>

<template>
  <div
    class="flex items-center justify-between p-3 bg-white rounded-lg shadow-sm hover:shadow-md transition-shadow"
  >
    <div class="flex items-center gap-3">
      <input
        type="checkbox"
        :checked="task.completed"
        @change="toggleCompleted"
        class="h-5 w-5 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
      />
      <span :class="{ 'line-through text-gray-500': task.completed }">
        {{ task.title }}
      </span>
    </div>
    <button
      @click="taskStore.deleteTask(task.id)"
      class="text-gray-400 hover:text-red-500"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M6 18L18 6M6 6l12 12"
        />
      </svg>
    </button>
  </div>
</template>
