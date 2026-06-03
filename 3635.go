package main

func earliestFinishTime2(lst []int, ldu []int, wst []int, wdu []int) int {
    ans := int(1e9)
    mln := int(1e9)
    mnw := int(1e9)

    for i := 0; i < len(lst); i++ {
        mln = min(mln, lst[i]+ldu[i])
    }

    for i := 0; i < len(wst); i++ {
        ans = min(ans, max(mln, wst[i])+wdu[i])
    }

    for i := 0; i < len(wst); i++ {
        mnw = min(mnw, wst[i]+wdu[i])
    }

    for i := 0; i < len(lst); i++ {
        ans = min(ans, max(mnw, lst[i])+ldu[i])
    }

    return ans
}