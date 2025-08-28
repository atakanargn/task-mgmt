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
    content: string;
    completed: boolean;
    board_id: string;
    created_at: string;
}

export interface Board {
    id: string;
    title: string;
    tasks: Task[];
}

export const useBoardStore = defineStore('board', {
    state: () => ({
        boards: [] as Board[],
        loading: false,
        selectedTask: null as Task | null,
    }),
    actions: {
        selectTask(task: Task) {
            this.selectedTask = task;
        },
        clearSelectedTask() {
            this.selectedTask = null;
        },
        async fetchBoards() {
            this.loading = true;
            try {
                const response = await apiClient.get('/boards');
                this.boards = response.data || [];
            } catch (error) {
                console.error('Failed to fetch boards:', error);
            } finally {
                this.loading = false;
            }
        },
        async addBoard(title: string) {
            if (!title) return;
            try {
                const response = await apiClient.post('/boards', { title });
                this.boards.push({ ...response.data, tasks: [] });
            } catch (error) {
                console.error('Failed to add board:', error);
            }
        },
        async addTask(boardId: string, title: string) {
            if (!title) return;
            try {
                const response = await apiClient.post('/tasks', { title, board_id: boardId });
                const board = this.boards.find(b => b.id === boardId);
                if (board) {
                    board.tasks.push(response.data);
                }
            } catch (error) {
                console.error('Failed to add task:', error);
            }
        },
        async updateTask(task: Task) {
            try {
                const response = await apiClient.patch(`/tasks/${task.id}`, {
                    title: task.title,
                    completed: task.completed,
                    board_id: task.board_id,
                    content: task.content,
                });

                const board = this.boards.find(b => b.id === task.board_id);
                if (board) {
                    const taskIndex = board.tasks.findIndex(t => t.id === task.id);
                    if (taskIndex !== -1) {
                        board.tasks[taskIndex] = response.data;
                    }
                }
            } catch (error) {
                console.error('Failed to update task:', error);
            }
        },
        async deleteTask(boardId: string, taskId: string) {
            try {
                await apiClient.delete(`/tasks/${taskId}`);
                const board = this.boards.find(b => b.id === boardId);
                if (board) {
                    board.tasks = board.tasks.filter(t => t.id !== taskId);
                }
            } catch (error) {
                console.error('Failed to delete task:', error);
            }
        }
    }
})
