import axios from 'axios';
import { defineStore } from 'pinia';

const apiClient = axios.create({
    baseURL: '/api',
    headers: {
        'Content-Type': 'application/json',
    },
});

export interface Task {
    id: string;
    title: string;
    completed: boolean;
    created_at: string;
}

export const useTaskStore = defineStore('task', {
    state: () => ({
        tasks: [] as Task[],
        loading: false,
    }),
    actions: {
        async fetchTasks() {
            this.loading = true;
            try {
                const response = await apiClient.get('/tasks');
                this.tasks = response.data || [];
            } catch (error) {
                console.error('Failed to fetch tasks:', error);
            } finally {
                this.loading = false;
            }
        },
        async addTask(title: string) {
            if (!title) return;
            try {
                const response = await apiClient.post('/tasks', { title });
                this.tasks.unshift(response.data);
            } catch (error) {
                console.error('Failed to add task:', error);
            }
        },
        async updateTask(task: Task) {
            try {
                await apiClient.patch(`/tasks/${task.id}`, {
                    title: task.title,
                    completed: task.completed
                });
            } catch (error) {
                console.error('Failed to update task:', error);
            }
        },
        async deleteTask(id: string) {
            try {
                await apiClient.delete(`/tasks/${id}`);
                this.tasks = this.tasks.filter(t => t.id !== id);
            } catch (error) {
                console.error('Failed to delete task:', error);
            }
        }
    }
})
