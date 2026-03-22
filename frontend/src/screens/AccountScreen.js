import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, ScrollView, Alert, SafeAreaView } from 'react-native';
import tw from 'twrnc';
import { useStore } from '../store/useStore';

export default function AccountScreen() {
  const { accounts, fetchAccounts } = useStore();
  const addAccount = async (name, type, pb) => {
    try {
      // Direct API call without action abstraction for brevity
      const api = (await import('../api/axios')).default;
      await api.post('/accounts', { 
        name, 
        type, 
        initial_balance: parseFloat(pb) || 0,
        currency: 'CNY'
      });
      fetchAccounts();
    } catch (e) {
      Alert.alert('Error', 'Failed to add account');
    }
  };

  const [name, setName] = useState('');
  const [type, setType] = useState('bank');
  const [balance, setBalance] = useState('');

  const handleCreate = () => {
    if (!name) return Alert.alert('Error', 'Name is required');
    addAccount(name, type, balance);
    setName('');
    setBalance('');
  };

  return (
    <SafeAreaView style={tw`flex-1 bg-gray-50`}>
      <ScrollView contentContainerStyle={tw`p-4`}>
        <Text style={tw`text-2xl font-bold text-gray-800 mb-6`}>资产账户</Text>
        
        {/* Account List */}
        {accounts.map(acc => (
          <View key={acc.id} style={tw`bg-white p-4 rounded-xl mb-3 shadow-sm border-l-4 border-blue-500 flex-row justify-between items-center`}>
            <View>
              <Text style={tw`text-lg font-bold text-gray-800`}>{acc.name}</Text>
              <Text style={tw`text-sm text-gray-500`}>{acc.type.toUpperCase()}</Text>
            </View>
            <Text style={tw`text-lg font-bold text-gray-800`}>¥ {acc.balance.toFixed(2)}</Text>
          </View>
        ))}

        {/* Add Account Form */}
        <View style={tw`mt-6 bg-white p-5 rounded-2xl shadow-sm`}>
          <Text style={tw`text-lg font-bold text-gray-800 mb-4`}>新增账户</Text>
          <TextInput 
            style={tw`bg-gray-100 p-3 rounded-lg mb-3 text-black`}
            placeholder="账户名称 (例: 招商银行)"
            value={name}
            onChangeText={setName}
          />
          <View style={tw`flex-row mb-3`}>
            {['bank', 'alipay', 'wechat', 'cash'].map(t => (
              <TouchableOpacity 
                key={t}
                onPress={() => setType(t)}
                style={tw`flex-1 py-2 mx-1 rounded-lg items-center ${type === t ? 'bg-blue-100 border border-blue-500' : 'bg-gray-100'}`}
              >
                <Text style={tw`text-xs ${type === t ? 'text-blue-600 font-bold' : 'text-gray-500'}`}>{t.toUpperCase()}</Text>
              </TouchableOpacity>
            ))}
          </View>
          <TextInput 
            style={tw`bg-gray-100 p-3 rounded-lg mb-4 text-black`}
            placeholder="初始余额 (例: 1000)"
            keyboardType="numeric"
            value={balance}
            onChangeText={setBalance}
          />
          <TouchableOpacity 
            style={tw`bg-blue-600 py-3 rounded-lg items-center`}
            onPress={handleCreate}
          >
            <Text style={tw`text-white font-bold`}>确认添加</Text>
          </TouchableOpacity>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}
