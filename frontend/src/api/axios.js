import axios from 'axios';
import AsyncStorage from '@react-native-async-storage/async-storage';

const api = axios.create({
  baseURL: 'http://10.0.2.2:8080/api/v1', // use localhost / 10.0.2.2 for Android emulator
  timeout: 10000,
});

api.interceptors.request.use(
  async (config) => {
    const token = await AsyncStorage.getItem('userToken');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

api.interceptors.response.use(
  (response) => {
    return response.data;
  },
  (error) => {
    // Check if error is unauthorized (401)
    if (error.response && error.response.status === 401) {
      // Handle logout logic or token refresh here
    }
    return Promise.reject(error);
  }
);

export default api;
