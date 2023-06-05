package service

import (
    "regexp"
)

// GetJiraID 从字符串中提取 Jira ID
func GetJiraID(s string) string {
    // 定义一个正则表达式，匹配 feature-XXX-YYY 格式的字符串
    re := regexp.MustCompile(`^feature-([A-Za-z0-9]+-\d+)$`)
    // 使用正则表达式匹配字符串
    match := re.FindStringSubmatch(s)
    // 如果匹配成功，返回第一个捕获组的内容
    if len(match) == 2 {
        return match[1]
    }
    // 如果匹配失败，返回空字符串
    return ""
}