<template>
  <div>
    <div class="flex border h-16 items-center">
      <Button variant="ghost" class="mx-6" @click="this.$router.push('/')">
        <h2 class="text-lg font-semibold">
          KubePanopticon
        </h2>
      </Button>
      <div class="ml-auto flex w-full space-x-2 justify-end mx-6">
        <Button @click="this.$router.push('/cluster')">Cluster</Button>
        <Button @click="this.$router.push('/nodes')">Node</Button>
        <Button>Application</Button>
        <Button>Metrics</Button>
      </div>
    </div>
  </div>
</template>

<script>
import Button from "@/components/ui/button/Button.vue";


export default {
  name: "NavBar",
  components: {Button},
  methods: {
    toNodePage() {
      this.$router.push('/nodes')
    }
  }
}
</script>


<style scoped>

</style>