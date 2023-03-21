export const mixins = {
  mounted() {
    this.provideActiveIndex()
  },
  methods: {
    // 导航栏激活标签传值给父组件
    provideActiveIndex() {
      this.$emit('activeIndex', this.activeIndex)
    },
  },
}
