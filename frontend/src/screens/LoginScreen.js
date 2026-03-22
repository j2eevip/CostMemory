import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, Alert, SafeAreaView } from 'react-native';
import tw from 'twrnc';
import { useStore } from '../store/useStore';

export default function LoginScreen({ navigation }) {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const login = useStore((state) => state.login);

  const handleLogin = async () => {
    try {
      await login(email, password);
    } catch (e) {
      Alert.alert('Error', 'Invalid credentials');
    }
  };

  return (
    <SafeAreaView style={tw`flex-1 bg-white justify-center items-center`}>
      <View style={tw`w-4/5`}>
        <Text style={tw`text-3xl font-bold text-center text-blue-600 mb-8`}>记账 App</Text>
        
        <TextInput 
          style={tw`bg-gray-100 p-4 rounded-xl mb-4 text-black`}
          placeholder="Email"
          autoCapitalize="none"
          keyboardType="email-address"
          value={email}
          onChangeText={setEmail}
        />
        
        <TextInput 
          style={tw`bg-gray-100 p-4 rounded-xl mb-6 text-black`}
          placeholder="Password"
          secureTextEntry
          value={password}
          onChangeText={setPassword}
        />

        <TouchableOpacity 
          style={tw`bg-blue-600 py-4 rounded-xl items-center`}
          onPress={handleLogin}
        >
          <Text style={tw`text-white font-bold text-lg`}>登录</Text>
        </TouchableOpacity>

        <TouchableOpacity 
          style={tw`mt-6 items-center`}
          onPress={() => navigation.navigate('Register')}
        >
          <Text style={tw`text-blue-500 font-semibold`}>没有账号？去注册</Text>
        </TouchableOpacity>
      </View>
    </SafeAreaView>
  );
}
