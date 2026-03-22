import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, ScrollView, Alert, SafeAreaView } from 'react-native';
import tw from 'twrnc';
import { useStore } from '../store/useStore';

export default function AddTransactionScreen({ navigation }) {
  const { accounts, categories, addTransaction } = useStore();

  const [amount, setAmount] = useState('');
  const [description, setDescription] = useState('');
  const [type, setType] = useState('expense'); // 'income' or 'expense'
  const [selectedAccount, setSelectedAccount] = useState(null);
  const [selectedCategory, setSelectedCategory] = useState(null);

  const handleSubmit = async () => {
    if (!amount || isNaN(amount)) return Alert.alert('错误', '请输入有效金额');
    if (!selectedAccount) return Alert.alert('错误', '请选择账户');
    if (!selectedCategory) return Alert.alert('错误', '请选择分类');

    try {
      await addTransaction({
        account_id: selectedAccount,
        category_id: selectedCategory,
        amount: parseFloat(amount),
        type,
        description,
        transaction_date: new Date().toISOString().split('T')[0],
      });
      Alert.alert('成功', '记账成功！', [
        { text: 'OK', onPress: () => navigation.navigate('Dashboard') }
      ]);
      setAmount('');
      setDescription('');
    } catch (e) {
      Alert.alert('错误', '记账失败，请重试');
    }
  };

  return (
    <SafeAreaView style={tw`flex-1 bg-gray-50`}>
      <ScrollView contentContainerStyle={tw`p-5`}>
        <Text style={tw`text-2xl font-bold text-gray-800 mb-6 text-center`}>记一笔</Text>
        
        {/* Type Switcher */}
        <View style={tw`flex-row bg-gray-200 p-1 rounded-xl mb-6`}>
          <TouchableOpacity 
            style={tw`flex-1 py-3 rounded-lg items-center ${type === 'expense' ? 'bg-white shadow' : ''}`}
            onPress={() => setType('expense')}
          >
            <Text style={tw`font-bold ${type === 'expense' ? 'text-red-500' : 'text-gray-500'}`}>支出</Text>
          </TouchableOpacity>
          <TouchableOpacity 
            style={tw`flex-1 py-3 rounded-lg items-center ${type === 'income' ? 'bg-white shadow' : ''}`}
            onPress={() => setType('income')}
          >
            <Text style={tw`font-bold ${type === 'income' ? 'text-green-500' : 'text-gray-500'}`}>收入</Text>
          </TouchableOpacity>
        </View>

        {/* Amount Input */}
        <View style={tw`bg-white p-6 rounded-2xl shadow-sm mb-4 flex-row items-center justify-between`}>
          <Text style={tw`text-3xl font-bold text-gray-800`}>¥</Text>
          <TextInput 
            style={tw`flex-1 text-right text-4xl text-gray-800 font-bold ml-2`}
            placeholder="0.00"
            keyboardType="numeric"
            value={amount}
            onChangeText={setAmount}
          />
        </View>

        <TextInput 
          style={tw`bg-white p-4 rounded-xl shadow-sm mb-6 text-gray-800`}
          placeholder="备注信息 (选填)"
          value={description}
          onChangeText={setDescription}
        />

        {/* Account Selection */}
        <Text style={tw`text-sm font-bold text-gray-500 mb-2 ml-1`}>选择账户</Text>
        <ScrollView horizontal showsHorizontalScrollIndicator={false} style={tw`mb-4`}>
          {accounts.map(acc => (
            <TouchableOpacity 
              key={acc.id}
              onPress={() => setSelectedAccount(acc.id)}
              style={tw`mr-3 px-5 py-3 rounded-xl border ${selectedAccount === acc.id ? 'bg-blue-50 border-blue-500' : 'bg-white border-transparent shadow-sm'}`}
            >
              <Text style={tw`font-semibold ${selectedAccount === acc.id ? 'text-blue-600' : 'text-gray-800'}`}>{acc.name}</Text>
            </TouchableOpacity>
          ))}
        </ScrollView>

        {/* Category Selection */}
        <Text style={tw`text-sm font-bold text-gray-500 mb-2 ml-1`}>选择分类</Text>
        <ScrollView horizontal showsHorizontalScrollIndicator={false} style={tw`mb-8`}>
          {categories.filter(c => c.type === type).map(cat => (
            <TouchableOpacity 
              key={cat.id}
              onPress={() => setSelectedCategory(cat.id)}
              style={tw`mr-3 px-5 py-3 rounded-xl border ${selectedCategory === cat.id ? 'bg-blue-50 border-blue-500' : 'bg-white border-transparent shadow-sm'}`}
            >
              <Text style={tw`font-semibold ${selectedCategory === cat.id ? 'text-blue-600' : 'text-gray-800'}`}>{cat.name}</Text>
            </TouchableOpacity>
          ))}
          {categories.filter(c => c.type === type).length === 0 && (
            <Text style={tw`text-gray-400 italic py-2`}>请先去设置页面添加分类</Text>
          )}
        </ScrollView>

        <TouchableOpacity 
          style={tw`bg-blue-600 py-4 rounded-xl items-center shadow-lg`}
          onPress={handleSubmit}
        >
          <Text style={tw`text-white font-bold text-lg`}>保存记录</Text>
        </TouchableOpacity>

      </ScrollView>
    </SafeAreaView>
  );
}
