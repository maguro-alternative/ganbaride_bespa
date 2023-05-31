package main

import "bonus/bonus"
import "fmt"
import "sort"

func main(){
	var zenei bonus.Aptitude
	var kouei bonus.Aptitude

	var enezenei bonus.Aptitude
	var enekouei bonus.Aptitude

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

	team := bonus.RiderTeam{
		Zenei			: zenei,
		Kouei			: kouei,
		Attack			: zenei.Attack,
		Defence			: zenei.Defence,
		HitPoint		: zenei.HitPoint + kouei.HitPoint,
		FinishingAttack	: zenei.FinishingAttack,
		AttackPoint		: teamAttackPoint,
	}

	teamBonus := bonus.TeamBonus{
		Attack			: 0,
		Defence			: 0,
		HitPoint		: 0,
		FinishingAttack	: 0,
	}

	teamSkill := bonus.NewSkillBonus()

	teamTeki := bonus.RiderTeamTekisei{
		AttackTeki: zenei.AttackTeki + kouei.AttackTeki,
		DefenceTeki: zenei.DefenceTeki + kouei.DefenceTeki,
		HitPointTeki: zenei.HitPointTeki + kouei.HitPointTeki,
		FinishingAttackTeki: zenei.FinishingAttackTeki + kouei.FinishingAttackTeki,
	}

	fields := []int{
		teamTeki.AttackTeki,
		teamTeki.DefenceTeki,
		teamTeki.HitPointTeki,
		teamTeki.FinishingAttackTeki,
	}

	sortFields := fields

	// スライスをソート
	sort.Sort(sort.Reverse(sort.IntSlice(sortFields)))

	bonusStetus1 := (sortFields[0] + 1) * 50
	bonusStetus2 := sortFields[1] * 50

	bonusStetus1Index := bonus.ContainsPoint(fields,sortFields[0])
	bonusStetus2Index := bonus.ContainsPoint(fields,sortFields[1])

	if (bonusStetus1Index == 0){
		teamBonus.Attack = bonusStetus1
	}else if (bonusStetus1Index == 1) {
		teamBonus.Defence = bonusStetus1
	}else if (bonusStetus1Index == 2) {
		teamBonus.HitPoint = bonusStetus1
	}else if (bonusStetus1Index == 3) {
		teamBonus.FinishingAttack = bonusStetus1
	}

	if (bonusStetus2Index == 0){
		teamBonus.Attack = bonusStetus2
	}else if (bonusStetus2Index == 1) {
		teamBonus.Defence = bonusStetus2
	}else if (bonusStetus2Index == 2) {
		teamBonus.HitPoint = bonusStetus2
	}else if (bonusStetus2Index == 3) {
		teamBonus.FinishingAttack = bonusStetus2
	}

	fmt.Println(teamBonus)

	// 相性適性の最高値「3」が一致し、合計値が「6」になると成立する。
	if (teamTeki.AttackTeki == 6 && 
		zenei.AttackTeki == kouei.AttackTeki) {
		fmt.Println("攻撃ベストパートナー")
	}
	if (teamTeki.DefenceTeki == 6 &&
		zenei.DefenceTeki == kouei.DefenceTeki) {
		fmt.Println("防御ベストパートナー")
	}
	if (teamTeki.HitPointTeki == 6 &&
		zenei.HitPointTeki == kouei.HitPointTeki) {
		fmt.Println("体力ベストパートナー")
	}
	if (teamTeki.FinishingAttackTeki == 6 &&
		zenei.FinishingAttackTeki == kouei.FinishingAttackTeki) {
		fmt.Println("必殺ベストパートナー")
	}

	zeneiSkillBonus := zenei.ZeneiRiderSkill(kouei,enezenei,enekouei)
	koueiSkillBonus := kouei.KoueiRiderSkill(zenei,enezenei,enekouei)

	if (zeneiSkillBonus.MyAttack != 0) {
		teamSkill.MyAttack += zeneiSkillBonus.MyAttack
	}
	if (zeneiSkillBonus.MyDefence != 0) {
		teamSkill.MyDefence += zeneiSkillBonus.MyDefence
	}
	if (zeneiSkillBonus.MyHitPoint != 0) {
		teamSkill.MyHitPoint += zeneiSkillBonus.MyHitPoint
	}
	if (zeneiSkillBonus.MyFinishingAttack != 0) {
		teamSkill.MyFinishingAttack += zeneiSkillBonus.MyFinishingAttack
	}
	if (zeneiSkillBonus.MyAttackPoint != 0) {
		teamSkill.MyAttackPoint += zeneiSkillBonus.MyAttackPoint
	}
	if (zeneiSkillBonus.EnemyAttack != 0) {
		teamSkill.EnemyAttack += zeneiSkillBonus.EnemyAttack
	}
	if (zeneiSkillBonus.EnemyDefence != 0) {
		teamSkill.EnemyDefence += zeneiSkillBonus.EnemyDefence
	}
	if (zeneiSkillBonus.EnemyHitPoint != 0) {
		teamSkill.EnemyHitPoint += zeneiSkillBonus.EnemyHitPoint
	}
	if (zeneiSkillBonus.EnemyFinishingAttack != 0) {
		teamSkill.EnemyFinishingAttack += zeneiSkillBonus.EnemyFinishingAttack
	}
	if (zeneiSkillBonus.EnemyAttackPoint != 0) {
		teamSkill.EnemyAttackPoint += zeneiSkillBonus.EnemyAttackPoint
	}
	

	if (koueiSkillBonus.MyAttack != 0) {
		teamSkill.MyAttack += koueiSkillBonus.MyAttack
	}
	if (koueiSkillBonus.MyDefence != 0) {
		teamSkill.MyDefence += koueiSkillBonus.MyDefence
	}
	if (koueiSkillBonus.MyHitPoint != 0) {
		teamSkill.MyHitPoint += koueiSkillBonus.MyHitPoint
	}
	if (koueiSkillBonus.MyFinishingAttack != 0) {
		teamSkill.MyFinishingAttack += koueiSkillBonus.MyFinishingAttack
	}
	if (koueiSkillBonus.MyAttackPoint != 0) {
		teamSkill.MyAttackPoint += koueiSkillBonus.MyAttackPoint
	}
	if (koueiSkillBonus.EnemyAttack != 0) {
		teamSkill.EnemyAttack += koueiSkillBonus.EnemyAttack
	}
	if (koueiSkillBonus.EnemyDefence != 0) {
		teamSkill.EnemyDefence += koueiSkillBonus.EnemyDefence
	}
	if (koueiSkillBonus.EnemyHitPoint != 0) {
		teamSkill.EnemyHitPoint += koueiSkillBonus.EnemyHitPoint
	}
	if (koueiSkillBonus.EnemyFinishingAttack != 0) {
		teamSkill.EnemyFinishingAttack += koueiSkillBonus.EnemyFinishingAttack
	}
	if (koueiSkillBonus.EnemyAttackPoint != 0) {
		teamSkill.EnemyAttackPoint += koueiSkillBonus.EnemyAttackPoint
	}

	fmt.Printf("コウゲキ:%d\n",team.Attack+teamBonus.Attack+teamSkill.MyAttack)
	fmt.Printf("ボウギョ:%d\n",team.Defence+teamBonus.Defence+teamSkill.MyDefence)
	fmt.Printf("タイリョク:%d\n",team.HitPoint+teamBonus.HitPoint+teamSkill.MyHitPoint)
	fmt.Printf("アタックポイント:%d\n",team.AttackPoint+teamSkill.MyAttackPoint)


	fmt.Println(team.Zenei.Name)
	fmt.Println(team.Kouei.Name)
	//fmt.Println(teamTeki)
	//fmt.Println(teamBonus)
	//fmt.Println(teamSkill)
	
}
