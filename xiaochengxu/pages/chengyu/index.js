//index.js
//获取应用实例
const app = getApp()
const fileManager = wx.getFileSystemManager();
Page({
  data: {
    motto: '笔画练字',
    userInfo: {},
    hasUserInfo: false,
    canIUse: wx.canIUse('button.open-type.getUserInfo'),
    pageData: {},
    hanzi:'',
    search:''
  },
  onShow() {
    if(app.globalData.hanzi) {
      this.getChengyuData();
      app.globalData.hanzi = ""
    }
    if (typeof this.getTabBar === 'function' &&
      this.getTabBar()) {
      this.getTabBar().setData({
        selected: 1
      })
    }
  },
  click_index(e){
    app.globalData.hanzi = e.currentTarget.dataset.zi
    wx.switchTab({
      url: '/pages/index/index'
    })
  },
  onLoad: function () {
    
    wx.showShareMenu({

      withShareTicket:true,
      
      menus:['shareAppMessage','shareTimeline']
      
    });
    this.getChengyuData();
    if (app.globalData.userInfo) {
      this.setData({
        userInfo: app.globalData.userInfo,
        hasUserInfo: true
      })
    } else if (this.data.canIUse){
      // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
      // 所以此处加入 callback 以防止这种情况
      app.userInfoReadyCallback = res => {
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        })
      }
    } else {
      // 在没有 open-type=getUserInfo 版本的兼容处理
      wx.getUserInfo({
        success: res => {
          app.globalData.userInfo = res.userInfo
          this.setData({
            userInfo: res.userInfo,
            hasUserInfo: true
          })
        }
      })
    }
  },
  getUserInfo: function(e) {
    app.globalData.userInfo = e.detail.userInfo
    this.setData({
      userInfo: e.detail.userInfo,
      hasUserInfo: true
    })
  },

  getChengyuData: function(e) {
    self = this;
    var search = self.data.search;
    if(app.globalData.hanzi) {
      search = app.globalData.hanzi
    }
    wx.request({
      url: '',
      data: {
        token:''
      },
      success: function(res) {
        var gif = [];
        var png = [];
        var hanzi = res.data.hanzi.split(",");
        for(var index = 0; index < res.data.hanzi_list.length; index++) {
          gif[index] = res.data.hanzi_list[index]['bihua_gif'];
          png[index] = res.data.hanzi_list[index]['bihua_png'];
        }
        res.data.hanzi_list = hanzi
        res.data.gif = gif
        res.data.png = png
        self.setData({pageData:res.data})
      },
      error: function(res) {
      }
    })
  },
  searchHanzi: function(e) {
    var hanzi = this.data.hanzi;
    if (hanzi == "") {
      return;
    }
    this.setData({search:hanzi})
    this.getChengyuData();
  },
  inputHanzi: function(e) {
    this.setData({hanzi:e.detail.value})
  }
})
