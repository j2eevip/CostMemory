import React, { useEffect } from 'react';
import { View, Text, ScrollView, TouchableOpacity, RefreshControl, SafeAreaView } from 'react-native';
import tw from 'twrnc';
import { useStore } from '../store/useStore';
import { LogOut, RefreshCcw } from 'lucide-react-native';

export default function DashboardScreen() {
  const { 
    accounts, 
    transactions, 
    fetchAccounts, 
    fetchTransactions, 
    fetchCategories,
    logout 
  } = useStore();

  const [refreshing, setRefreshing] = React.useState(false);

  const loadData = async () => {
    setRefreshing(true);
    await Promise.all([
      fetchAccounts(),
      fetchTransactions(),
      fetchCategories()
    ]);
    setRefreshing(false);
  };

  useEffect(() => {
    loadData();
  }, []);

  const totalBalance = accounts.reduce((sum, acc) => sum + acc.balance, 0).toFixed(2);
  const recentTransactions = transactions.slice(0, 5);

  return (
    <SafeAreaView style={tw`flex-1 bg-gray-50`}>
      <ScrollView
        contentContainerStyle={tw`p-4`}
        refreshControl={<RefreshControl refreshing={refreshing} onRefresh={loadData} />}
      >
        <View style={tw`flex-row justify-between items-center mb-6`}>
          <Text style={tw`text-2xl font-bold text-gray-800`}>概览</Text>
          <TouchableOpacity onPress={logout} style={tw`p-2 bg-red-100 rounded-full`}>
            <LogOut size={20} color="#ef4444" />
          </TouchableOpacity>
        </View>

        {/* Balance Card */}
        <View style={tw`bg-blue-600 rounded-2xl p-6 shadow-md mb-6`}>
          <Text style={tw`text-blue-100 text-sm mb-1`}>总净资产 (CNY)</Text>
          <Text style={tw`text-white text-4xl font-extrabold`}>¥ {totalBalance}</Text>
        </View>

        {/* Recent Transactions */}
        <View style={tw`mb-4`}>
          <Text style={tw`text-lg font-bold text-gray-800 mb-3`}>最近交易</Text>
          {recentTransactions.length === 0 ? (
            <Text style={tw`text-gray-500 italic`}>暂无交易记录</Text>
          ) : (
            recentTransactions.map(tx => (
              <View key={tx.id} style={tw`flex-row justify-between items-center bg-white p-4 rounded-xl mb-2 shadow-sm`}>
                <View style={tw`flex-row items-center`}>
                  <View style={tw`w-10 h-10 rounded-full justify-center items-center mr-3 ${tx.type === 'expense' ? 'bg-red-100' : 'bg-green-100'}`}>
                    <Text style={tw`text-lg`}>{tx.type === 'expense' ? '💸' : '💰'}</Text>
                  </View>
                  <View>
                    <Text style={tw`text-base font-bold text-gray-800`}>{tx.description || '无备注'}</Text>
                    <Text style={tw`text-xs text-gray-500`}>{new Date(tx.transaction_date).toLocaleDateString()}</Text>
                  </View>
                </View>
                <Text style={tw`text-base font-bold ${tx.type === 'expense' ? 'text-red-500' : 'text-green-500'}`}>
                  {tx.type === 'expense' ? '-' : '+'}¥{tx.amount.toFixed(2)}
                </Text>
              </View>
            ))
          )}
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}
