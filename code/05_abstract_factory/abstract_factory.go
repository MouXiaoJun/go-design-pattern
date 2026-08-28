package _5_abstract_factory

// Button 抽象产品：按钮
type Button interface {
	Render() string
}

// Text 抽象产品：文本
type Text interface {
	Render() string
}

// ThemeFactory 抽象工厂：创建一族风格一致的 UI 组件
type ThemeFactory interface {
	CreateButton() Button
	CreateText() Text
}

// DarkThemeFactory 暗色主题工厂
type DarkThemeFactory struct{}

func (DarkThemeFactory) CreateButton() Button { return &darkButton{} }
func (DarkThemeFactory) CreateText() Text     { return &darkText{} }

type darkButton struct{}

func (*darkButton) Render() string {
	return `<button style="background:#222;color:#eee">确定</button>`
}

type darkText struct{}

func (*darkText) Render() string { return `<span style="color:#eee">暗色文本</span>` }

// LightThemeFactory 浅色主题工厂
type LightThemeFactory struct{}

func (LightThemeFactory) CreateButton() Button { return &lightButton{} }
func (LightThemeFactory) CreateText() Text     { return &lightText{} }

type lightButton struct{}

func (*lightButton) Render() string {
	return `<button style="background:#fff;color:#222">确定</button>`
}

type lightText struct{}

func (*lightText) Render() string { return `<span style="color:#222">浅色文本</span>` }

// RenderUI 使用抽象工厂渲染整套 UI，客户端只依赖接口，不关心具体主题
func RenderUI(f ThemeFactory) string {
	return f.CreateButton().Render() + "\n" + f.CreateText().Render()
}
