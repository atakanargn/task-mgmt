<script setup lang="ts">
import { onMounted } from "vue";
import TaskForm from "./components/TaskForm.vue";
import TaskItem from "./components/TaskItem.vue";
import { useTaskStore } from "./stores/taskStore";

const taskStore = useTaskStore();

onMounted(() => {
  taskStore.fetchTasks();
});
</script>

<template>
  <div class="bg-gray-100 min-h-screen">
    <div class="max-w-2xl mx-auto py-10 px-4">
      <header class="text-center mb-8">
        <h1 class="text-4xl font-bold text-gray-800">Task Management</h1>
      </header>

      <main class="bg-white p-6 rounded-xl shadow-lg">
        <TaskForm />
        <div v-if="taskStore.loading" class="text-center text-gray-500">
          Loading...
        </div>
        <div v-else class="space-y-3">
          <TaskItem
            v-for="task in taskStore.tasks"
            :key="task.id"
            :task="task"
          />
        </div>
      </main>
    </div>
  </div>
</template>
