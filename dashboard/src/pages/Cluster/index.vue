<template>
  <div>
    <div class="flex flex-col justify-center">
      <div class="w-full">
        <div id="cpu-usage" class="w-1/2 h-80"></div>
      </div>
    </div>
  </div>
</template>

<script>
import * as echarts from 'echarts';
import {getClusterAllPodCPUUsage} from "@/api/module/cluster.js";
import Button from "@/components/ui/button/Button.vue";

export default {
  name: "ClusterBase",
  components: {Button},

  data() {
    return {
      cluster_all_pods_cpu_usage: []
    }
  },

  methods: {
    formatTime(timestamp) {
      const date = new Date(timestamp * 1000); // JavaScript的时间戳是毫秒
      return date.toLocaleString();
    },
    getClusterAllPodsCPUUsage(){
      const myChart = echarts.init(document.getElementById('cpu-usage'));

      getClusterAllPodCPUUsage().then(response => {
        const data = response.data.data[0].values.map(item => ({
          value: [this.formatTime(item[0]), item[1]]
        }));
        myChart.setOption({
          xAxis: {
            type: 'category',
            boundaryGap: false,
          },
          yAxis: {
            type: 'value'
          },
          dataZoom: [
            {
              type: 'inside',
              start: 0,
              end: 20
            },
            {
              start: 0,
              end: 20
            }
          ],
          series: [{
            data: data,
            type: 'line',
            smooth: true, // 使线条平滑
            areaStyle: {
              opacity: 0.8 // 设置填充区域的透明度
            },
            lineStyle: {
              width: 1, // 设置线条宽度，可以设置为0隐藏折线
              opacity: 0.5 // 设置线条透明度，可以调整以淡化折线
            },
            itemStyle: {
              opacity: 0 // 将节点透明度设置为0，以隐藏节点
            },
            symbol: 'none' // 不显示折线图上的节点
          }]
        });
      })

    }
  },

  mounted() {
    this.getClusterAllPodsCPUUsage();
  }
}
</script>

<style scoped>

</style>