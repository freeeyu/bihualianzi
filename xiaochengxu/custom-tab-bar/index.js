Component({
  data: {
    selected: 0,
    // color: "#7A7E83",
    // selectedColor: "#3cc51f",
    list: [{
      pagePath: "/pages/index/index",
      iconPath: "/image/word1.png",
      selectedIconPath: "/image/word2.png",
      text: "练字"
    }, {
      pagePath: "/pages/chengyu/index",
      iconPath: "/image/chengyu1.png",
      selectedIconPath: "/image/chengyu2.png",
      text: "成语"
    }]
  },
  attached() {
  },
  methods: {
    switchTab(e) {
      const data = e.currentTarget.dataset
      const url = data.path
      wx.switchTab({url})
      this.setData({
        selected: data.index
      })
    }
  }
})