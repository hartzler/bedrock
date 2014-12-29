#!/usr/bin/env ruby

#
# helper functions for determining optimal shard count for smooth scaling of number of nodes
#


# find the divisors of n
def div(n); (2..n-1).map{|i| n%i==0 ? i : nil}.compact; end

# show the numbers >= n with the most divisors
def top(n); (1..n).map{|i| [div(i).size,i]}.sort.reverse[0..10]; end

# show numbers <= m which include n as a divisor
def y(n,m=4096); (1..m).map{|i| div(i).include?(n) ? i : nil}.compact; end

# show the number of shards per node given m*r shards (m shards, r is replicas)
def shards(m, r=1); div(m).map{|n| [n,m/n*r]}; end

# [60,120,360,720,1680,2520].map{|n| div(n).size}
# => [10, 14, 22, 28, 38, 46]
