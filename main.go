package main

import "bonus/bonus"
import "fmt"

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
	fmt.Println(team.Attack)
	fmt.Println(zenei.ZeneiRiderSkill(kouei,enezenei,enekouei))
	
}
