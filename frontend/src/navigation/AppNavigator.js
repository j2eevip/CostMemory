import React from 'react';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { NavigationContainer } from '@react-navigation/native';
import { Compass, PlusCircle, User, CreditCard } from 'lucide-react-native';

import LoginScreen from '../screens/LoginScreen';
import RegisterScreen from '../screens/RegisterScreen';
import DashboardScreen from '../screens/DashboardScreen';
import AddTransactionScreen from '../screens/AddTransactionScreen';
import AccountScreen from '../screens/AccountScreen';
import { useStore } from '../store/useStore';

const Stack = createNativeStackNavigator();
const Tab = createBottomTabNavigator();

function MainTabs() {
  return (
    <Tab.Navigator screenOptions={{ headerShown: false, tabBarActiveTintColor: '#3b82f6' }}>
      <Tab.Screen 
        name="Dashboard" 
        component={DashboardScreen} 
        options={{ tabBarIcon: ({ color }) => <Compass color={color} size={24} /> }} 
      />
      <Tab.Screen 
        name="Add" 
        component={AddTransactionScreen} 
        options={{ tabBarIcon: ({ color }) => <PlusCircle color={color} size={32} /> }} 
      />
      <Tab.Screen 
        name="Accounts" 
        component={AccountScreen} 
        options={{ tabBarIcon: ({ color }) => <CreditCard color={color} size={24} /> }} 
      />
    </Tab.Navigator>
  );
}

export default function AppNavigator() {
  const userToken = useStore((state) => state.userToken);

  return (
    <NavigationContainer>
      <Stack.Navigator screenOptions={{ headerShown: false }}>
        {userToken == null ? (
          <>
            <Stack.Screen name="Login" component={LoginScreen} />
            <Stack.Screen name="Register" component={RegisterScreen} />
          </>
        ) : (
          <Stack.Screen name="MainTabs" component={MainTabs} />
        )}
      </Stack.Navigator>
    </NavigationContainer>
  );
}
