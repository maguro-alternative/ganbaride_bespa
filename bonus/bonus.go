package bonus

import (
	"fmt"
)
/*
ライダースキルの構造体
*/
type SkillBonus struct {
	MyAttack				int
	MyDefence 				int
	MyHitPoint 				int
	MyFinishingAttack 		int
	MyAttackPoint			int
	EnemyAttack				int
	EnemyDefence 			int
	EnemyHitPoint 			int
	EnemyFinishingAttack 	int
	EnemyAttackPoint		int
}
// コンストラクタ関数を定義する
func NewSkillBonus() *SkillBonus {
	// デフォルト値を設定する
	return &SkillBonus{
	  MyAttack: 0,
	  MyDefence: 0,
	  MyHitPoint: 0,
	  MyFinishingAttack: 0,
	  MyAttackPoint: 0,
	  EnemyAttack: 0,
	  EnemyDefence: 0,
	  EnemyHitPoint: 0,
	  EnemyFinishingAttack: 0,
	  EnemyAttackPoint: 0,
	}
}

// チームボーナス
type TeamBonus struct {
	Attack				int
	Defence 			int
	HitPoint 			int
	FinishingAttack		int
}

// ライダーステータス
type Aptitude struct {
	Name				string
	Category			[]string
	Attribute			string
	Attack 				int
	Defence 			int
	HitPoint 			int
	FinishingAttack 	int
	SkillName			string
	SkillType			int
	SkillConditions		string
	SkillBonus			SkillBonus
	AttackTeki 			int
	DefenceTeki 		int
	HitPointTeki 		int
	FinishingAttackTeki int
}

/*
0	mujouken
1	zenei
2	kouei
3	zenei no toki nakama zokusei
4	kouei no toki nakama zokusei
5	nakama zokusei
6	zenei no toki nakama rider(cate[])
7	kouei no toki nakama rider(cate[])
8	aite zokusei
9	zenei no toki aite zokusei
10	kouei no toki aite zokusei
11	aite rider[cate]
12	zenei no toki aite rider(cate[])
13	kouei no toki aite rider(cate[])
*/
// func (変数名 構造体) 関数名(変数名 型) 型
func (zenei Aptitude) ZeneiRiderSkill(
	kouei Aptitude,
	eneZen Aptitude,
	eneKou Aptitude,
) SkillBonus {
	var zeneiSkillType int = zenei.SkillType
	var bonus SkillBonus = *NewSkillBonus()
	
	// 無条件 ゼンエイの場合
	if (zeneiSkillType == 0 || zeneiSkillType == 1) {
		bonus = zenei.SkillBonus
	// ゼンエイ時、仲間と属性が一緒の場合
	}else if (zeneiSkillType == 3 && kouei.Attribute == zenei.Attribute){
		bonus = zenei.SkillBonus
	// 仲間と属性が一緒の場合
	}else if (zeneiSkillType == 5 && kouei.Attribute == zenei.Attribute){
		bonus = zenei.SkillBonus
	// ゼンエイ時、仲間が○○の場合
	}else if (zeneiSkillType == 6 && contains(kouei.Category,zenei.SkillConditions)){
		bonus = zenei.SkillBonus
	// 相手と属性が同じ場合
	}else if (zeneiSkillType == 8 && zenei.Attribute == eneZen.Attribute){
		bonus = zenei.SkillBonus
	// ゼンエイ時、相手と属性が同じ場合
	}else if (zeneiSkillType == 9 && zenei.Attribute == eneZen.Attribute){
		bonus = zenei.SkillBonus
	// 相手が○○の場合
	}else if (zeneiSkillType == 11 && (contains(eneZen.Category,zenei.SkillConditions) || contains(eneKou.Category,zenei.SkillConditions))){
		bonus = zenei.SkillBonus
	// ゼンエイジ相手に○○がいる場合
	}else if (zeneiSkillType == 12 && (contains(eneZen.Category,zenei.SkillConditions) || contains(eneKou.Category,zenei.SkillConditions))){
		bonus = zenei.SkillBonus
	}
	return bonus
}

func (kouei Aptitude) KoueiRiderSkill(
	zenei Aptitude,
	eneZen Aptitude,
	eneKou Aptitude,
) SkillBonus {
	var koueiSkillType int = kouei.SkillType
	var bonus SkillBonus = *NewSkillBonus()
	
	// 無条件 コウエイの場合
	if (koueiSkillType == 0 || koueiSkillType == 2) {
		bonus = kouei.SkillBonus
	// コウエイ時、仲間と属性が一緒の場合
	}else if (koueiSkillType == 4 && 
		kouei.Attribute == zenei.Attribute){

		bonus = kouei.SkillBonus
	// 仲間と属性が一緒の場合
	}else if (koueiSkillType == 5 && 
		kouei.Attribute == zenei.Attribute){

		bonus = kouei.SkillBonus
	// コウエイ時、仲間が○○の場合
	}else if (koueiSkillType == 7 && contains(
		zenei.Category,kouei.SkillConditions)){

		bonus = kouei.SkillBonus
	// 相手と属性が同じ場合
	}else if (koueiSkillType == 8 && 
		kouei.Attribute == eneZen.Attribute){

		bonus = kouei.SkillBonus
	// コウエイ時、相手と属性が同じ場合
	}else if (koueiSkillType == 10 && 
		kouei.Attribute == eneZen.Attribute){

		bonus = kouei.SkillBonus
	// 相手が○○の場合
	}else if (koueiSkillType == 11 && (
		contains(eneZen.Category,kouei.SkillConditions) || 
		contains(eneKou.Category,kouei.SkillConditions))){

		bonus = kouei.SkillBonus
	// コウエイ時相手に○○がいる場合
	}else if (koueiSkillType == 13 && (
		contains(eneZen.Category,kouei.SkillConditions) || 
		contains(eneKou.Category,kouei.SkillConditions))){

		bonus = kouei.SkillBonus
	}
	return bonus
}

// チームステータス
type RiderTeam struct {
	Zenei			Aptitude
	Kouei			Aptitude
	Attack 			int
	Defence 		int
	HitPoint 		int
	FinishingAttack int
	AttackPoint		int
}

type RiderTeamTekisei struct {
	AttackTeki 			int
	DefenceTeki 		int
	HitPointTeki 		int
	FinishingAttackTeki int
}

func main(){
	var zenei Aptitude
	var kouei Aptitude

	var teamAttackPoint int = 0


	// ゼンエイステータス
	// 5-001
	// 仮面ライダーディケイド コンプリートフォーム
	zenei.Name = "仮面ライダーディケイド コンプリートフォーム"
	zenei.Category = []string{"ディケイド"}
	zenei.Attribute = "技"
	zenei.Attack = 450
	zenei.Defence = 400
	zenei.HitPoint = 550
	zenei.FinishingAttack = 2600
	zenei.SkillName = "最終形態完成"
	zenei.SkillType = 1
	zenei.SkillBonus.MyDefence = 300
	zenei.AttackTeki = 3
	zenei.DefenceTeki = 1
	zenei.HitPointTeki = 0
	zenei.FinishingAttackTeki = 2

	if (zenei.Attribute == "速"){
		teamAttackPoint = 10
	}

	// コウエイステータス
	// 8-053
	// シャドームーン
	kouei.Name = "シャドームーン"
	kouei.Category = []string{"シャドームーン"}
	kouei.Attribute = "力"
	kouei.Attack = 450
	kouei.Defence = 400
	kouei.HitPoint = 500
	kouei.FinishingAttack = 2350
	kouei.SkillName = "天・海・地"
	kouei.SkillType = 0
	kouei.SkillBonus.MyAttack = 150
	kouei.SkillBonus.MyDefence = 150
	kouei.SkillBonus.MyHitPoint = 150
	kouei.AttackTeki = 3
	kouei.DefenceTeki = 1
	kouei.HitPointTeki = 1
	kouei.FinishingAttackTeki = 0

	team := RiderTeam{
		Zenei			: zenei,
		Kouei			: kouei,
		Attack			: zenei.Attack,
		Defence			: zenei.Defence,
		HitPoint		: zenei.HitPoint + kouei.HitPoint,
		FinishingAttack	: zenei.FinishingAttack,
		AttackPoint		: teamAttackPoint,
	}

	// 相性適性の最高値「3」が一致し、合計値が「6」になると成立する。
	if (zenei.AttackTeki + kouei.AttackTeki == 6 && 
		zenei.AttackTeki == kouei.AttackTeki) {
		fmt.Println("攻撃ベストパートナー")
	}
	if (zenei.DefenceTeki + kouei.DefenceTeki == 6 &&
		zenei.DefenceTeki == kouei.DefenceTeki) {
		fmt.Println("防御ベストパートナー")
	}
	if (zenei.HitPointTeki + kouei.HitPointTeki == 6 &&
		zenei.HitPointTeki == kouei.HitPointTeki) {
		fmt.Println("体力ベストパートナー")
	}
	if (zenei.FinishingAttackTeki + kouei.FinishingAttackTeki == 6 &&
		zenei.FinishingAttackTeki == kouei.FinishingAttackTeki) {
		fmt.Println("必殺ベストパートナー")
	}
	fmt.Println(team)
	
}


// 配列に指定したものが含まれているか確認
func contains(elems []string, v string) bool {
    for _, s := range elems {
        if v == s {
            return true
        }
    }
    return false
}

// 配列に指定したものが含まれていた場合、場所を返す
func ContainsPoint(elems []int, v int) int {
	var i int = 0
    for _, s := range elems {
        if v == s {
            return i
        }
		i += 1
    }
    return -1
}