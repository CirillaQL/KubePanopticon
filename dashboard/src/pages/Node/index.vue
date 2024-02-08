<template>
  <div class="flex justify-center">
    <div class="overflow-x-auto w-3/4 ">
      <Table>
        <TableCaption>节点信息</TableCaption>
        <TableHeader>
          <TableHead>Name</TableHead>
          <TableHead>Status</TableHead>
        </TableHeader>
        <TableBody>
          <TableRow v-for="(nodeItem, index) in nodesList" :key="'nodeItem/'+index">
            <TableCell>
              {{ nodeItem.node.metadata.name }}
            </TableCell>
            <TableCell>
              <div v-for="(status, statusIndex) in nodeItem.status" :key="'nodeItemStatus/'+statusIndex">
                <Badge class="bg-green-500 m-0.5" v-if="status === 'Ready'">{{ status }}</Badge>
                <Badge variant="destructive" class="m-0.5" v-else>{{ status }}</Badge>
              </div>
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>
  </div>
</template>

<script>
import {getNodeList} from "@/api/module/nodes.js";

import {Table} from '@/components/ui/table'
import TableHead from "@/components/ui/table/TableHead.vue";
import TableCaption from "@/components/ui/table/TableCaption.vue";
import TableHeader from "@/components/ui/table/TableHeader.vue";
import TableCell from "@/components/ui/table/TableCell.vue";
import TableBody from "@/components/ui/table/TableBody.vue";
import TableRow from "@/components/ui/table/TableRow.vue";
import Badge from "@/components/ui/badge/Badge.vue";

export default {
  name: "NodeBase",
  components: {Badge, TableRow, TableBody, TableCell, Table, TableHeader, TableCaption, TableHead},

  data() {
    return {
      nodesList: [],
      nodesListLoading: false,
    }
  },

  methods: {
    getNodeList() {
      this.nodesListLoading = true
      getNodeList().then(response => {
        if (response.data.code === 0){
          this.nodesList = response.data.data
        }
        this.nodesListLoading = false
      })
    },
    // getNodeStatus(conditions) {
    //
    // }
  },

  mounted() {
    this.getNodeList()
  }
}
</script>

<style scoped>

</style>