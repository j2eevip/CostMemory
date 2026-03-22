import { create } from 'zustand';
import { persist, createJSONStorage } from 'zustand/middleware';
import AsyncStorage from '@react-native-async-storage/async-storage';
import api from '../api/axios';

export const useStore = create(
  persist(
    (set, get) => ({
      userToken: null,
      user: null,
      accounts: [],
      categories: [],
      transactions: [],

      // Auth actions
      login: async (email, password) => {
        try {
          const response = await api.post('/auth/login', { email, password });
          const token = response.token;
          await AsyncStorage.setItem('userToken', token);
          set({ userToken: token });
        } catch (error) {
          throw error;
        }
      },
      register: async (username, email, password) => {
        try {
          await api.post('/auth/register', { username, email, password });
        } catch (error) {
          throw error;
        }
      },
      logout: async () => {
        await AsyncStorage.removeItem('userToken');
        set({ userToken: null, user: null });
      },

      // Data fetching
      fetchAccounts: async () => {
        try {
          const res = await api.get('/accounts');
          set({ accounts: res.data });
        } catch (error) {
          console.error('Failed to fetch accounts', error);
        }
      },
      fetchCategories: async () => {
        try {
          const res = await api.get('/categories');
          set({ categories: res.data });
        } catch (error) {
          console.error('Failed to fetch categories', error);
        }
      },
      fetchTransactions: async () => {
        try {
          const res = await api.get('/transactions');
          set({ transactions: res.data });
        } catch (error) {
          console.error('Failed to fetch transactions', error);
        }
      },
      addTransaction: async (data) => {
        try {
          await api.post('/transactions', data);
          get().fetchTransactions();
          get().fetchAccounts(); // refresh balance
        } catch (error) {
          console.error('Failed to add transaction', error);
          throw error;
        }
      }
    }),
    {
      name: 'finance-app-storage',
      storage: createJSONStorage(() => AsyncStorage),
      partialize: (state) => ({ userToken: state.userToken }), // Persist only userToken for now
    }
  )
);
