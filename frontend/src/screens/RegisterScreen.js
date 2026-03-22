import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, Alert, SafeAreaView } from 'react-native';
import tw from 'twrnc';
import { useStore } from '../store/useStore';

export default function RegisterScreen({ navigation }) {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const register = useStore((state) => state.register);

  const handleRegister = async () => {
    try {
      if (!username || !email || !password) {
        Alert.alert('Error', 'Please fill in all fields');
        return;
      }
      await register(username, email, password);
      Alert.alert('Success', 'Registered! You can now log in.', [
        { text: 'OK', onPress: () => navigation.navigate('Login') }
      ]);
    } catch (e) {
      Alert.alert('Error', 'Registration failed');
    }
  };

  return (
    <SafeAreaView style={tw`flex-1 bg-white justify-center items-center`}>
      <View style={tw`w-4/5`}>
        <Text style={tw`text-3xl font-bold text-center text-blue-600 mb-8`}>注册</Text>
        
        <TextInput 
          style={tw`bg-gray-100 p-4 rounded-xl mb-4 text-black`}
          placeholder="Username"
          autoCapitalize="none"
          value={username}
          onChangeText={setUsername}
        />

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
          onPress={handleRegister}
        >
          <Text style={tw`text-white font-bold text-lg`}>注册并继续</Text>
        </TouchableOpacity>

        <TouchableOpacity 
          style={tw`mt-6 items-center`}
          onPress={() => navigation.goBack()}
        >
          <Text style={tw`text-gray-500 font-semibold`}>返回登录</Text>
        </TouchableOpacity>
      </View>
    </SafeAreaView>
  );
}
