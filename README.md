# Bluemountain
[toc]

原版请看[Markdown Reference (typora.com.cn)](https://support.typora.com.cn/Markdown-Reference/)



# 标题1

## 标题2

### 标题3

##### 标题4

>>>
>>>
>>>> > > 
>
>>>
>>
>>>
>
>​                



## un-ordered list

* Red
* Green
* Blue

## ordered list 

1. red
2. Green
3. Blue

- [ ] use Typora

```c
#include<stdio.h>
int main()
{
	printf("hello,world");
}
```

$$ \mathbf{V}_1 \times \mathbf{V}_2 =  \begin{vmatrix} \mathbf{i} & \mathbf{j} & \mathbf{k} \\ \frac{\partial X}{\partial u} &  \frac{\partial Y}{\partial u} & 0 \\ \frac{\partial X}{\partial v} &  \frac{\partial Y}{\partial v} & 0 \\ \end{vmatrix} $$


| 我是伞兵 | 我是麻痹 | 我是括约肌 |
| :------------ |:---------------:| :----:|
| 属于是 | 破防了 | 绝绝子 |
| XXX，但是XXX | 蚌埠住了 | 鼠鼠 |

小夫我要进来了[^1]

[^1]:出自禾野男孩

***

---
[是男人就来砍我](http://www.snszg.cn/ "传奇霸业私服")



跳转到[标题1](#标题1)ctrl+鼠标左键点击



我是[4399快乐器][id] 

[id]: http://www.4399.com/	"“好康的”"





[Baidu][] And then define the link: 

[Baidu]: http://Baidu.com/



<www.google.com>



<img src="D:\Typora\2022\企划\新年Typora开启计划\QQ图片20210921121827.jpg" style="zoom: 15%;" />

### 对字符的改变

*wish<em>guis*

_ws<em>wda_

[猫片](file://S:/heritage/大学生活)，注意要使用英文符号

**woshisha**

**bi**

__wshishabi__

<!----注释 ---->

` import this`

~~miakeken text~~

<u>fuck</u>

awd



 ==\=\= \=\===

<span style="color:red; background-color:pink">this text is red</span>

:clown_face:   :joy:    :point_up_2:  :dancers:

### LATEX

$\lim_{x \to \infty} \exp(-x) = 0$


$$
\ce{H2}+\ce{O2}\xlongequal[]{王艳}\ce{H2O}
$$

$$
H_2 O
$$

$$
H^2 O
$$





<iframe height='265' scrolling='no' title='Fancy Animated SVG Menu' src='http://codepen.io/jeangontijo/embed/OxVywj/?height=265&theme-id=0&default-tab=css,result&embed-version=2' frameborder='no' allowtransparency='true' allowfullscreen='true' style='width: 100%;'> </iframe>

使用 <kbd>Ctrl</kbd>+<kbd>Alt</kbd>+<kbd>Del</kbd> 重启电脑
<kbd> </kbd> -- 白色框框



## Mermaid

```mermaid
graph TD;
	/-->bin;
	/-->boot;
	/-->dev;
	/-->etc;
	/-->home;
	/-->root;
	/-->run;
	/-->sbin;
	/-->tmp;
	/-->usr;
	/-->var;
	home-->alice;
	home-->bob;
	home-->eve;
```

```mermaid
sequenceDiagram
    Alice->>+John: Hello John, how are you?
    Alice->>+John: John, can you hear me?
    John-->>-Alice: Hi Alice, I can hear you!
    John-->>-Alice: I feel great!

```

```mermaid
classDiagram
      Animal <|-- Duck
      Animal <|-- Fish
      Animal <|-- Zebra
      Animal : +int age
      Animal : +String gender
      Animal: +isMammal()
      Animal: +mate()
      class Duck{
          +String beakColor
          +swim()
          +quack()
      }
      class Fish{
          -int sizeInFeet
          -canEat()
      }
      class Zebra{
          +bool is_wild
          +run()
      }
```



```mermaid
gantt
    title 三下乡
    dateFormat DD
    section 廖青松
    A task           :a1, 2022-07-07, 5d
    Another task     :after a1  , 1d
    section 李豪
    Task in sec      :2022-07-07  , 5d
     section 数模三人
    Task in sec      :2022-07-07  , 5d
    another task      : 1d
     section 杨欢、陈鑫、王海霞、代宇、李玉梅
    Task in sec      :2022-07-07  , 5d
```

```mermaid
pie
    title Key elements in Product X
    "Calcium" : 42.96
    "Potassium" : 50.05
    "Magnesium" : 10.01
    "Iron" :  5
```



```mermaid
graph TD
    A[Christmas] -->|Get money| B(Go shopping)
    B --> C{Let me think}
    C -->|One| D[Laptop]
    C -->|Two| E[iPhone]
    C -->|Three| F[fa:fa-car Car]
```

```mermaid
stateDiagram-v2
    [*] --> Still
    Still --> [*]
    Still --> Moving
    Moving --> Still
    Moving --> Crash
    Crash --> [*]
            
```

```mermaid
erDiagram
          CUSTOMER }|..|{ DELIVERY-ADDRESS : has
          CUSTOMER ||--o{ ORDER : places
          CUSTOMER ||--o{ INVOICE : "liable for"
          DELIVERY-ADDRESS ||--o{ ORDER : receives
          INVOICE ||--|{ ORDER : covers
          ORDER ||--|{ ORDER-ITEM : includes
          PRODUCT-CATEGORY ||--|{ PRODUCT : contains
          PRODUCT ||--o{ ORDER-ITEM : "ordered in"
            
```

```mermaid
  journey
    title My working day
    section Go to work
      Make tea: 5: Me
      Go upstairs: 3: Me
      Do work: 1: Me, Cat
    section Go home
      Go downstairs: 5: Me
      Sit down: 3: Me
```

```mermaid
gitGraph:
options
{
    "nodeSpacing": 150,
    "nodeRadius": 10
}
end
commit
branch newbranch
checkout newbranch
commit
branch develop
checkout develop
commit
checkout master
commit
merge newbranch
```



```go
package main

import "fmt"

func main() {
    fmt.Println("Hello world") 
}

```
